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
) (
	map[string]string,
	error,
) {
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

func (auth *FirebaseAuthService) VerifyEmail(
	c echo.Context,
	payload *requests.VerifyEmailRequest,
) (map[string]string, error) {
	user := models.User{}
	result := auth.db.Where("email = ?", payload.Email).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, echo.NewHTTPError(
			http.StatusNotFound,
			errors.New("User not found with the specified email."),
		)
	}

	ident, err := utils.VerifyOTP(c, payload.Otp)
	if err != nil {
		return nil, err
	}

	if ident != nil {
		now := time.Now()
		user.EmailVerifiedAt = &now
		auth.db.Save(&user)
	}

	return map[string]string{"message": "User email successfully verified."}, nil
}

func (auth *FirebaseAuthService) Login(
	c echo.Context,
	payload *requests.LoginRequest,
) (
	map[string]string,
	error,
) {
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

	var userAvatar string
	if user.Avatar != nil {
	    userAvatar = *user.Avatar
	}

	return map[string]string{
		"uid": user.UID,
		"name": user.Name,
		"avatar": userAvatar,
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
