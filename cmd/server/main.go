package main

import (
	"github.com/Sahil2k07/gRPC-GO/internal/config"
	"github.com/Sahil2k07/gRPC-GO/internal/database"
	"github.com/Sahil2k07/gRPC-GO/internal/handler"
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	configs := config.LoadConfig()
	database.Connect()

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     configs.Origins,
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))

	public := e.Group("public")
	handler.HandlePublicEndpoints(public)

	secure := e.Group("api/v1")
	secure.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(configs.JWT.Secret),
		TokenLookup: "cookie:" + configs.JWT.CookieName,
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid or expired token"})
		},
	}))
	handler.HandleSecureEndpoints(secure)

	e.Logger.Fatal(e.Start(":5000"))
}
