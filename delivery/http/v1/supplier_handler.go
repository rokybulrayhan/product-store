package v1

import (
	"fmt"
	"strconv"

	"github.com/techno/entity/httpentity"
	"github.com/techno/lib/logger"
	Supplier "github.com/techno/service/supplier"

	"github.com/labstack/echo/v4"
)

type SupplierHandler struct {
	Services *Supplier.Service
	Logger   logger.Logger
}

func NewSupplierHandler(services *Supplier.Service, logger logger.Logger) *SupplierHandler {
	return &SupplierHandler{
		Services: services,
		Logger:   logger,
	}
}

// Map Suppliers routes
func (h SupplierHandler) MapSupplierRoutes(SupplierGroup *echo.Group, authenticated echo.MiddlewareFunc) {
	SupplierGroup.POST("", h.AddSupplierHandler)
	SupplierGroup.PUT("/:id", h.UpdateSupplierHandler)
	SupplierGroup.GET("/:id", h.GetSupplierHandler)
	SupplierGroup.GET("", h.ListSupplierHandler)
	SupplierGroup.DELETE("/:id", h.Delete)

}

func (h *SupplierHandler) ListSupplierHandler(c echo.Context) error {
	params := httpentity.SupplierParams{}
	err := c.Bind(&params)
	if err != nil {
		return handleInvalidDataError(c, err)
	}
	res, err := h.Services.List(c.Request().Context(), params)
	return handleApplicationResponse(c, "", res, err)
}

func (h *SupplierHandler) UpdateSupplierHandler(c echo.Context) error {

	SupplierId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return handleApplicationError(c, err)
	}
	data := httpentity.UpdateSupplierRequest{}
	data.Id = SupplierId

	// Bind JSON to struct
	err = c.Bind(&data)
	if err != nil {
		return handleInvalidDataError(c, err)
	}
	// Validate JSON data
	validationErrors := data.Validate()
	if validationErrors != nil {
		return handleValidationError(c, validationErrors)
	}
	res, err := h.Services.Update(c.Request().Context(), data)
	return handleApplicationResponse(c, "", res, err)
}

func (h SupplierHandler) GetSupplierHandler(c echo.Context) error {

	SupplierId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return handleApplicationError(c, err)
	}
	res, err := h.Services.GetByID(c.Request().Context(), SupplierId)
	return handleApplicationResponse(c, "", res, err)
}

func (h SupplierHandler) AddSupplierHandler(c echo.Context) error {
	data := httpentity.CreateSupplierRequest{}
	fmt.Println(data)

	err := c.Bind(&data)
	if err != nil {
		return handleInvalidDataError(c, err)
	}
	validationErrors := data.Validate()
	if validationErrors != nil {
		return handleValidationError(c, validationErrors)
	}

	res, err := h.Services.Create(c.Request().Context(), data)
	return handleApplicationResponse(c, "", res, err)
}

func (h *SupplierHandler) Delete(c echo.Context) error {
	addressIdStr := c.Param("id")
	addressId, err := strconv.Atoi(addressIdStr)
	if err != nil {
		return handleApplicationError(c, err)
	}

	ctx := c.Request().Context()
	err = h.Services.Delete(ctx, addressId)
	return handleApplicationResponse(c, "deleted successfully.", nil, err)
}
