package handler

import (
	"github.com/Sahil2k07/gRPC-GO/internal/repository"
	"github.com/Sahil2k07/gRPC-GO/internal/service"

	"github.com/labstack/echo/v4"
)

func HandleSecureEndpoints(g *echo.Group) {
	userRepo := repository.NewUserRepository()

	userService := service.NewUserService(userRepo)

	NewUserHandler(g, userService)
}
