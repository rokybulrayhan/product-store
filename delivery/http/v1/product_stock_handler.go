package v1

import (
	"fmt"
	"strconv"

	"github.com/techno/entity/httpentity"
	"github.com/techno/lib/logger"
	ProductStock "github.com/techno/service/productstock"

	"github.com/labstack/echo/v4"
)

type ProductStockHandler struct {
	Services *ProductStock.Service
	Logger   logger.Logger
}

func NewProductStockHandler(services *ProductStock.Service, logger logger.Logger) *ProductStockHandler {
	return &ProductStockHandler{
		Services: services,
		Logger:   logger,
	}
}

// Map ProductStocks routes
func (h ProductStockHandler) MapProductStockRoutes(ProductStockGroup *echo.Group, authenticated echo.MiddlewareFunc) {
	ProductStockGroup.POST("", h.AddProductStockHandler)
	ProductStockGroup.PUT("/:id", h.UpdateProductStockHandler)
	ProductStockGroup.GET("/:id", h.GetProductStockHandler)
	ProductStockGroup.GET("", h.ListProductStockHandler)
	ProductStockGroup.DELETE("/:id", h.Delete)

}

func (h *ProductStockHandler) ListProductStockHandler(c echo.Context) error {
	params := httpentity.ProductStockParams{}
	err := c.Bind(&params)
	if err != nil {
		return handleInvalidDataError(c, err)
	}
	res, err := h.Services.List(c.Request().Context(), params)
	return handleApplicationResponse(c, "", res, err)
}

func (h *ProductStockHandler) UpdateProductStockHandler(c echo.Context) error {

	ProductStockId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return handleApplicationError(c, err)
	}
	data := httpentity.UpdateProductStockRequest{}
	data.Id = ProductStockId

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

func (h ProductStockHandler) GetProductStockHandler(c echo.Context) error {

	ProductStockId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return handleApplicationError(c, err)
	}
	res, err := h.Services.GetByID(c.Request().Context(), ProductStockId)
	return handleApplicationResponse(c, "", res, err)
}

func (h ProductStockHandler) AddProductStockHandler(c echo.Context) error {
	data := httpentity.CreateProductStockRequest{}
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

func (h *ProductStockHandler) Delete(c echo.Context) error {
	addressIdStr := c.Param("id")
	addressId, err := strconv.Atoi(addressIdStr)
	if err != nil {
		return handleApplicationError(c, err)
	}

	ctx := c.Request().Context()
	err = h.Services.Delete(ctx, addressId)
	return handleApplicationResponse(c, "deleted successfully.", nil, err)
}
