package handler

import (
	"github.com/Sahil2k07/gRPC-GO/internal/repository"
	"github.com/Sahil2k07/gRPC-GO/internal/service"

	"github.com/labstack/echo/v4"
)

func HandleSecureEndpoints(g *echo.Group) {
	userRepo := repository.NewUserRepository()
	inventoryGroupRepo := repository.NewInventoryGroupRepository()

	userService := service.NewUserService(userRepo)
	inventoryGroupService := service.NewInventoryGroupService(inventoryGroupRepo)

	NewUserHandler(g, userService)
	NewInventoryGroupHandler(g, inventoryGroupService)
}
