package utils

import (
	"encoding/json"
	"go-api/config"
	"math/rand"
	"net/http"
	"time"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/coocood/freecache"
)

type Otp struct {
	Retries		uint `json:"retries"`
	ExpiresAt	time.Time `json:"expires_at"`
	VerifiedAt	*time.Time `json:"verified_at"`
}

func SendOTP(
	c echo.Context,
	identifier string,
	otp_value string,
	callback func(string),
) error {
	cache := c.Get("cache").(*freecache.Cache)
	key := []byte(identifier)
	value := []byte(otp_value)
	otp_string, _ := cache.Get(key)

	expirySecs := config.GetConfig().Otp.ExpirySecs
	retrySecs := config.GetConfig().Otp.RetrySecs
	retries := config.GetConfig().Otp.Retries

	if len(otp_string) == 0 {
		otp := Otp{
			Retries: 0,
			ExpiresAt: time.Now().Add(time.Second * time.Duration(expirySecs)),
			VerifiedAt: nil,
		}

		otp_string, _ = json.Marshal(otp)
		if err := cache.Set(key, otp_string, int(retrySecs)); err != nil {
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				err,
			)
		}

		if err := cache.Set(value, key, int(expirySecs)); err != nil {
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				err,
			)
		}

		callback(otp_value)

		return nil
	}

	otp := Otp{}
	err := json.Unmarshal(otp_string, &otp)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			err,
		)
	}

	otp_expiration_diff := int(otp.ExpiresAt.Sub(time.Now()).Seconds())

	if otp.Retries >= retries {
		cool_down_time := otp.ExpiresAt.Add(-time.Second * time.Duration(expirySecs)).Add(time.Second * time.Duration(retrySecs))
		otp_cool_down_diff := int(cool_down_time.Sub(time.Now()).Seconds())

		error_string := fmt.Sprintf(
			"Too many OTP's requested, try again after: %d",
			otp_cool_down_diff,
		)

		return echo.NewHTTPError(
			http.StatusBadRequest,
			error_string,
		)
	}

	if otp_expiration_diff > 0 {
		error_string := fmt.Sprintf(
			"Recently OTP requested, try again after: %d",
			otp_expiration_diff,
		)

		return echo.NewHTTPError(
			http.StatusBadRequest,
			error_string,
		)
	}

	otp.Retries += 1
	otp.ExpiresAt = time.Now().Add(time.Second * time.Duration(expirySecs))
	otp.VerifiedAt = nil

	otp_string, _ = json.Marshal(otp)
	if err := cache.Set(key, otp_string, int(retrySecs)); err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			err,
		)
	}

	if err := cache.Set(value, key, int(expirySecs)); err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			err,
		)
	}

	callback(otp_value)

	return nil
}

func VerifyOTP(c echo.Context, value string) (bool, error) {
	return false, nil
}

func GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	otp := rand.Intn(900000) + 100000
	return fmt.Sprintf("%06d", otp)
}
