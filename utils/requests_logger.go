package utils

import (
	"os"
	"fmt"
	"time"
	"go-api/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupHTTPRequestsLogger(
	app *echo.Echo,
	logFilePath string,
	errorFilePath string,
) error {
	appConfig := config.GetConfig()

	logFile, err := os.OpenFile(
		logFilePath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		return err
	}

	errorFile, err := os.OpenFile(
		errorFilePath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		return err
	}

	logFunc := func(c echo.Context, v middleware.RequestLoggerValues) error {
		now := time.Now().UTC()

		if v.Error != nil {
			errorMessage := fmt.Sprintf("[%s] request-id:\"%s\", error: \"%s\"", now, v.RequestID, v.Error.Error())

			if appConfig.Debug {
				fmt.Print(errorMessage)
			} else {
				if _, err := errorFile.WriteString(errorMessage); err != nil {
					return err
				}
			}
		}
			
		logMessage := fmt.Sprintf("[%s] debug:\"%t\", request-id:\"%s\", ip:\"%s\", method:\"%s\", host:\"%s\", uri:\"%s\", status:\"%d\", execution-time:\"%s\", headers:%s\n", now, appConfig.Debug, v.RequestID, v.RemoteIP, v.Method, v.Host, v.URI, v.Status, v.Latency, v.Headers)

		if appConfig.Debug {
			fmt.Print(logMessage)
		} else {
			if _, err := logFile.WriteString(logMessage); err != nil {
				return err
			}
		}
		return nil
	}

	app.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig {
		LogLatency: true,
		LogRemoteIP: true,
		LogHost: true,
		LogMethod: true,
		LogURI: true,
		LogRequestID: true,
		LogStatus: true,
		LogContentLength: true,
		LogHeaders: []string {
			"Content-Length",
			"Content-Type",
			"Accept",
			"Authorization",
			"User-Agent",
		},
		LogValuesFunc: logFunc,
	}))

	return nil
}
