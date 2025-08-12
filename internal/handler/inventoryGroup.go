package handler

import (
	"net/http"

	"github.com/Sahil2k07/gRPC-GO/internal/auth"
	"github.com/Sahil2k07/gRPC-GO/internal/enum"
	interfaces "github.com/Sahil2k07/gRPC-GO/internal/interface"
	"github.com/Sahil2k07/gRPC-GO/internal/util"
	"github.com/Sahil2k07/gRPC-GO/internal/view"
	"github.com/labstack/echo/v4"
)

type inventoryGroupHandler struct {
	s interfaces.InventoryGroupService
}

func NewInventoryGroupHandler(g *echo.Group, s interfaces.InventoryGroupService) *inventoryGroupHandler {
	h := &inventoryGroupHandler{s: s}

	g.GET("/inventory/group/:id", h.GetInventoryGroup, auth.WithRole(enum.ADMIN))
	g.POST("/inventory/group", h.AddInventoryGroup)
	g.POST("/inventory/group/list", h.ListInventoryGroup)
	g.PUT("/inventory/group", h.UpdateInventoryGroup)
	g.DELETE("/inventory/group/:id", h.DeleteInventoryGroup)

	return h
}

func (h *inventoryGroupHandler) GetInventoryGroup(c echo.Context) error {
	id := c.Param("id")

	g, err := h.s.GetInventoryGroup(id)
	if err != nil {
		return util.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, g)
}

func (h *inventoryGroupHandler) AddInventoryGroup(c echo.Context) error {
	var req view.AddInventoryGroup

	if err := util.BindAndValidate(c, &req); err != nil {
		return util.HandleError(c, err)
	}

	if err := h.s.AddInventoryGroup(req); err != nil {
		return util.HandleError(c, err)
	}

	return c.NoContent(http.StatusOK)
}

func (h *inventoryGroupHandler) UpdateInventoryGroup(c echo.Context) error {
	var req view.InventoryGroupView

	if err := util.BindAndValidate(c, &req); err != nil {
		return util.HandleError(c, err)
	}

	if err := h.s.UpdateInventoryGroup(req); err != nil {
		return util.HandleError(c, err)
	}

	return c.NoContent(http.StatusOK)
}

func (h *inventoryGroupHandler) DeleteInventoryGroup(c echo.Context) error {
	id := c.Param("id")

	if err := h.s.DeleteInventoryGroup(id); err != nil {
		return util.HandleError(c, err)
	}

	return c.NoContent(http.StatusOK)
}

func (h *inventoryGroupHandler) ListInventoryGroup(c echo.Context) error {
	var req view.ListInventoryGroup

	if err := util.BindAndValidate(c, &req); err != nil {
		return util.HandleError(c, err)
	}

	list, err := h.s.ListInventoryGroup(req)
	if err != nil {
		return util.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, list)
}
