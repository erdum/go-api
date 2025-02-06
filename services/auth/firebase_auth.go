package auth

import (
	"go-api/requests"
	"go-api/config"
	"go-api/models"
	"net/http"
	"context"
	"reflect"
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

func (auth *FirebaseAuthService) Register(*requests.RegisterRequest) (
	map[string]string,
	error,
) {
	return nil, nil
}

func (auth *FirebaseAuthService) Login(*requests.LoginRequest) (
	map[string]string,
	error,
) {
	return nil, nil
}

func (auth *FirebaseAuthService) SignOn(
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
