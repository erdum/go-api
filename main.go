package main

import (
	"go-api/config"
	"go-api/controllers"
	"go-api/middlewares"
	"go-api/models"
	"go-api/services"
	"go-api/services/auth"
	"go-api/requests"
	"go-api/utils"
	"go-api/validators"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/go-faker/faker/v4"
)

func initialMigration() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{
		SkipDefaultTransaction: true,
		TranslateError: true,
	})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.UserAddress{},
		&models.TraineeDetail{},
		&models.CoachDetail{},
		&models.Session{},
		&models.SessionRequest{},
		&models.Booking{},
		&models.Media{},
		&models.CoachVerificationRequest{},
		&models.Review{},
		&models.Product{},
		&models.Cart{},
		&models.Order{},
		&models.DeliveryAddress{},
		&models.PaymentMethod{},
		&models.Bank{},
		&models.Withdrawal{},
		&models.Notification{},
	); err != nil {
		return nil, err
	}

	return db, nil
}

func dbSeed(db *gorm.DB, models ...interface{}) error {
	for _, model := range models {
		err := faker.FakeData(model)
		if err != nil {
			return err
		}

		db.Save(model)
    }

    return nil
}

func main() {
	app := echo.New()
	app.Use(middleware.RequestID())
	app.Use(middlewares.Cache())
	app.Validator = validators.NewDefaultValidator()

	appConfig, err := config.LoadConfig()
	if err != nil {
		app.Logger.Fatal(err)
	}

	db, err := initialMigration()
	if err != nil {
		app.Logger.Fatal(err)
	}

	// err = dbSeed(db, &models.User{})
	// if err != nil {
	// 	app.Logger.Fatal(err)
	// }

	// HTTP Requests log file
	err = utils.SetupHTTPRequestsLogger(app, "requests.log", "error.log")
	if err != nil {
		app.Logger.Fatal(err)
	}

	// Services
	tokenService := auth.NewJWTToken()
	authService := auth.NewFirebaseAuth(db, tokenService)
	userService := services.NewUserService(db)

	// Inject services into the controllers
	authController := controllers.NewAuthController(authService, tokenService)
	userController := controllers.NewUserController(userService)

	// app.GET("/users", userController.GetAllUsers)
	// app.POST("/users", userController.CreateUser)
	// app.GET("/users/:id", userController.GetUser)
	// app.PUT("/users/:id", userController.UpdateUser)
	// app.DELETE("/users/:id", userController.DeleteUser)

	app.POST(
		"/register",
		authController.Register,
		middlewares.Validate(&requests.RegisterRequest{}),
		middlewares.Throttle(4, 10),
	)

	app.POST(
		"/login",
		authController.Login,
		middlewares.Validate(&requests.LoginRequest{}),
	)

	app.POST(
		"/sign-on",
		authController.SignOn,
		middlewares.Validate(&requests.SignOnRequest{}),
	)

	app.POST(
		"/forget-password",
		authController.ForgetPassword,
		middlewares.Validate(&requests.ResendOtpRequest{}),
	)

	app.POST(
		"/update-password",
		authController.UpdatePassword,
		middlewares.Validate(&requests.UpdatePasswordRequest{}),
	)

	app.POST(
		"/resend-otp",
		authController.ResendOtp,
		middlewares.Validate(&requests.ResendOtpRequest{}),
	)

	app.POST(
		"/verify-otp",
		authController.VerifyOtp,
		middlewares.Validate(&requests.VerifyOtpRequest{}),
	)

	// Protected Routes (Require authentication)
	protectedRoutes := app.Group("")
	protectedRoutes.Use(middlewares.Authenticate(tokenService, db))

	// Update profile
	protectedRoutes.POST(
		"/profile",
		userController.UpdateProfile,
		middlewares.Validate(&requests.UpdateProfileRequest{}),
	)

	// Get profile
	protectedRoutes.GET(
		"/profile",
		userController.GetProfile,
	)

	// Set user location
	protectedRoutes.POST(
		"/user/location",
		userController.UpdateLocation,
		middlewares.Validate(&requests.UpdateLocationRequest{}),
	)

	app.Logger.Fatal(app.Start(":" + appConfig.Port))
}
