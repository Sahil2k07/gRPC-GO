package handler

import (
	"github.com/Sahil2k07/gRPC-GO/internal/repository"
	"github.com/Sahil2k07/gRPC-GO/internal/service"

	"github.com/labstack/echo/v4"
)

func HandlePublicEndpoints(g *echo.Group) {
	authRepo := repository.NewAuthRepository()

	authService := service.NewAuthService(authRepo)

	NewAuthHandler(g, authService)
}
