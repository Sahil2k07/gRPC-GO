package handler

import (
	"net/http"

	interfaces "github.com/Sahil2k07/gRPC-GO/internal/interface"
	"github.com/Sahil2k07/gRPC-GO/internal/util"
	"github.com/Sahil2k07/gRPC-GO/internal/view"
	"github.com/labstack/echo/v4"
)

type inventoryItemHandler struct {
	s interfaces.InventoryItemService
}

func NewInventoryItemHandler(g *echo.Group, s interfaces.InventoryItemService) *inventoryItemHandler {
	h := &inventoryItemHandler{s: s}

	g.GET("/inventory/item/:id", h.GetInventoryItem)
	g.POST("/inventory/item/list", h.ListInventoryItems)
	g.POST("/inventory/item", h.AddInventoryItem)
	g.PUT("/inventory/item", h.UpdateInventoryItem)
	g.DELETE("/inventory/item/:id", h.DeleteInventoryItem)
	g.PATCH("/inventory/item/stock", h.UpdateInventoryItemStock)

	return h
}

func (h *inventoryItemHandler) GetInventoryItem(c echo.Context) error {
	id := c.Param("id")

	record, err := h.s.GetInventoryItem(id)
	if err != nil {
		return util.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, record)
}

func (h *inventoryItemHandler) AddInventoryItem(c echo.Context) error {
	var req view.AddInventoryItem

	if err := util.BindAndValidate(c, &req); err != nil {
		return util.HandleError(c, err)
	}
	if err := h.s.AddInventoryItem(req); err != nil {
		return util.HandleError(c, err)
	}

	return c.NoContent(http.StatusOK)
}

func (h *inventoryItemHandler) UpdateInventoryItem(c echo.Context) error {
	var req view.UpdateInventoryItem

	if err := util.BindAndValidate(c, &req); err != nil {
		return util.HandleError(c, err)
	}

	if err := h.s.UpdateInventoryItem(req); err != nil {
		return util.HandleError(c, err)
	}

	return c.NoContent(http.StatusOK)
}

func (h *inventoryItemHandler) DeleteInventoryItem(c echo.Context) error {
	id := c.Param("id")

	if err := h.s.DeleteInventoryItem(id); err != nil {
		return util.HandleError(c, err)
	}

	return c.NoContent(http.StatusOK)
}

func (h *inventoryItemHandler) ListInventoryItems(c echo.Context) error {
	var req view.ListInventoryItem

	if err := util.BindAndValidate(c, &req); err != nil {
		return util.HandleError(c, err)
	}

	records, err := h.s.ListInventoryItems(req)
	if err != nil {
		return util.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, records)
}

func (h *inventoryItemHandler) UpdateInventoryItemStock(c echo.Context) error {
	var req []view.UpdateInventoryStock

	if err := util.BindAndValidate(c, &req); err != nil {
		return util.HandleError(c, err)
	}

	if err := h.s.UpdateInventoryItemStock(req); err != nil {
		return util.HandleError(c, err)
	}

	return c.NoContent(http.StatusOK)
}
