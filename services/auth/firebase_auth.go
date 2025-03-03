package auth

import (
	"go-api/requests"
	"go-api/config"
	"go-api/models"
	"go-api/utils"
	"net/http"
	"context"
	"reflect"
	"errors"
	"time"
	"fmt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"firebase.google.com/go"
	"google.golang.org/api/option"
)

type FirebaseAuthService struct {
	db *gorm.DB
	tokenService TokenService
	appConfig *config.Config
}

func NewFirebaseAuth(
	db *gorm.DB,
	tokenService TokenService,
) *FirebaseAuthService {
	return &FirebaseAuthService{
		db: db,
		tokenService: tokenService,
		appConfig: config.GetConfig(),
	}
}

func (auth *FirebaseAuthService) Register(
	c echo.Context,
	payload *requests.RegisterRequest,
) (map[string]string, error) {
	user := models.User{}
	result := auth.db.Where("email = ?", payload.Email).First(&user)

	if user.EmailVerifiedAt != nil {
		return nil, echo.NewHTTPError(
			http.StatusBadRequest,
			errors.New("Email already exists."),
		)
	}

	if result.RowsAffected > 0 {
		user.UID = utils.GenerateHexUUID()
		user.Name = payload.Name
		user.Email = payload.Email
		user.PhoneNumber = payload.PhoneNumber
		user.Password, _ = utils.HashPassword(payload.Password)
		auth.db.Save(&user)
	} else {
		user.UID = utils.GenerateHexUUID()
		user.Name = payload.Name
		user.Email = payload.Email
		user.PhoneNumber = payload.PhoneNumber
		user.Password, _ = utils.HashPassword(payload.Password)
		result = auth.db.Create(&user)

		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, echo.NewHTTPError(
				http.StatusBadRequest,
				errors.New("Phone Number already in use."),
			)
		}
	}

	// Send OTP
	err := utils.SendOTP(
		c,
		user.Email,
		utils.GenerateOTP(),
		func (value string) {
			fmt.Println("Email Sent.", value)
			subject := "OTP | "+auth.appConfig.Name
			content := "Hello "+user.Name+",\n\nHere is your verification code:\n\n"+value+"\n\nPlease use this code to complete your action.\n\nThank you,\n"+auth.appConfig.Name
			utils.SendMail(subject, content, []string{user.Email})
		},
	)

	if err != nil {
		return nil, err
	}

	return map[string]string{"message": "User successfully registered."}, nil
}

func (auth *FirebaseAuthService) Login(
	c echo.Context,
	payload *requests.LoginRequest,
) (map[string]string, error) {
	user := models.User{}
	result := auth.db.Where("email = ?", payload.Email).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, echo.NewHTTPError(
			// http.StatusNotFound,
			http.StatusBadRequest,
			// errors.New("User not found with the specified email."),
			errors.New("Invalid credentials."),
		)
	}

	if !utils.VerifyPassword(payload.Password, user.Password) {
		return nil, echo.NewHTTPError(
			http.StatusBadRequest,
			// errors.New("Invalid password."),
			errors.New("Invalid credentials."),
		)
	}

	if user.EmailVerifiedAt == nil {
		return nil, echo.NewHTTPError(
			http.StatusBadRequest,
			errors.New("Account email not verified."),
		)
	}

	if payload.FcmToken != "" {
		user.FcmToken = &payload.FcmToken
		auth.db.Save(&user)
	}

	token, err := auth.tokenService.GenerateToken(
		Token{
			EntityID: user.ID,
			EntityType: fmt.Sprintf("%T", reflect.TypeOf(user)),
		},
	)

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var userAvatar string
	if user.Avatar != nil {
	    userAvatar = *user.Avatar
	}

	return map[string]string{
		"uid": user.UID,
		"name": user.Name,
		"avatar": userAvatar,
		"token": token,
	}, nil
}

func (auth *FirebaseAuthService) SignOn(
	c echo.Context,
	payload *requests.SignOnRequest,
) (map[string]string, error) {
	ctx := context.Background()
	conf := &firebase.Config{
	    ProjectID: auth.appConfig.Firebase.ProjectId,
	}
	opt := option.WithCredentialsFile(auth.appConfig.Firebase.Credentials)
	app, err := firebase.NewApp(ctx, conf, opt)

	if err != nil {
	    return nil, echo.NewHTTPError(http.StatusBadRequest, err)
	}
	client, err := app.Auth(ctx)

	if err != nil {
	    return nil, echo.NewHTTPError(http.StatusBadRequest, err)
	}
	data, err := client.VerifyIDToken(ctx, payload.IdToken)

	if err != nil {
	    return nil, echo.NewHTTPError(http.StatusBadRequest, err)
	}

	userData := map[string]string{
		"email": data.Claims["email"].(string),
		"uid": data.Claims["user_id"].(string),
		"name": data.Claims["name"].(string),
		"avatar": data.Claims["picture"].(string),
	}
	userAvatar := userData["avatar"]

	user := models.User{}
	auth.db.Where(models.User{Email: userData["email"]}).Assign(
		models.User{
			Name: userData["name"],
			UID: userData["uid"],
			Avatar: &userAvatar,
		},
	).FirstOrCreate(&user)

	token, err := auth.tokenService.GenerateToken(
		Token{
			EntityID: user.ID,
			EntityType: fmt.Sprintf("%T", reflect.TypeOf(user)),
		},
	)

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return map[string]string{"token": token}, nil
}

func (auth *FirebaseAuthService) ForgetPassword(
	c echo.Context,
	payload *requests.ResendOtpRequest,
) (map[string]string, error) {
	user := models.User{}
	result := auth.db.Where("email = ?", payload.Email).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, echo.NewHTTPError(
			http.StatusNotFound,
			errors.New("User not found with the specified email."),
		)
	}

	now := time.Now()
	user.PasswordResetRequested = &now
	auth.db.Save(&user)

	_, err := auth.ResendOtp(c, payload)

	if err != nil {
		return nil, err
	}

	return map[string]string{
		"message": "Forget password successfully requested.",
	}, nil
}

func (auth *FirebaseAuthService) UpdatePassword(
	c echo.Context,
	payload *requests.UpdatePasswordRequest,
) (map[string]string, error) {
	user := models.User{}
	result := auth.db.Where("email = ?", payload.Email).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, echo.NewHTTPError(
			http.StatusNotFound,
			errors.New("User not found with the specified email."),
		)
	}

	var passwordResetRequested time.Time = time.Now()
	if user.PasswordResetRequested != nil {
		passwordResetRequested = *user.PasswordResetRequested
	}

	if passwordResetRequested.Before(time.Now()) {
		return nil, echo.NewHTTPError(
			http.StatusForbidden,
			errors.New(
				"Password update request expired, please request again.",
			),
		)
	}

	user.Password, _ = utils.HashPassword(payload.NewPassword)
	user.PasswordResetRequested = nil
	auth.db.Save(&user)

	return map[string]string{
		"message": "User password has been successfully updated",
	}, nil
}

func (auth *FirebaseAuthService) ResendOtp(
	c echo.Context,
	payload *requests.ResendOtpRequest,
) (map[string]string, error) {
	user := models.User{}
	result := auth.db.Where("email = ?", payload.Email).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, echo.NewHTTPError(
			http.StatusNotFound,
			errors.New("User not found with the specified email."),
		)
	}

	// Send OTP
	err := utils.SendOTP(
		c,
		user.Email,
		utils.GenerateOTP(),
		func (value string) {
			fmt.Println("Email Sent.", value)
			subject := "OTP | "+auth.appConfig.Name
			content := "Hello "+user.Name+",\n\nHere is your verification code:\n\n"+value+"\n\nPlease use this code to complete your action.\n\nThank you,\n"+auth.appConfig.Name
			utils.SendMail(subject, content, []string{user.Email})
		},
	)

	if err != nil {
		return nil, err
	}

	return map[string]string{"message": "OTP successfully resent."}, nil
}

func (auth *FirebaseAuthService) VerifyOtp(
	c echo.Context,
	payload *requests.VerifyOtpRequest,
) (map[string]string, error) {
	user := models.User{}
	result := auth.db.Where("email = ?", payload.Email).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, echo.NewHTTPError(
			http.StatusNotFound,
			errors.New("User not found with the specified email."),
		)
	}

	_, err := utils.VerifyOTP(c, payload.Otp, true)
	if err != nil {
		return nil, err
	}

	if user.EmailVerifiedAt == nil {
		now := time.Now()
		user.EmailVerifiedAt = &now
		auth.db.Save(&user)

		return map[string]string{
			"message": "User email successfully verified.",
		}, nil
	}

	if user.PasswordResetRequested != nil {
		expiry := time.Now().Add(time.Second * time.Duration(auth.appConfig.PasswordResetExpirySecs))
		user.PasswordResetRequested = &expiry
		auth.db.Save(&user)
	}

	return map[string]string{
		"message": "User OTP successfully verified.",
	}, nil
}
