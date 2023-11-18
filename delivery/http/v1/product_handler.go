package v1

import (
	"fmt"
	"strconv"

	"github.com/techno/entity/httpentity"
	"github.com/techno/lib/logger"
	Product "github.com/techno/service/product"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	Services *Product.Service
	Logger   logger.Logger
}

func NewProductHandler(services *Product.Service, logger logger.Logger) *ProductHandler {
	return &ProductHandler{
		Services: services,
		Logger:   logger,
	}
}

// Map Products routes
func (h ProductHandler) MapProductRoutes(ProductGroup *echo.Group, authenticated echo.MiddlewareFunc) {
	ProductGroup.POST("", h.AddProductHandler)
	ProductGroup.PUT("/:id", h.UpdateProductHandler)
	ProductGroup.GET("/:id", h.GetProductHandler)
	ProductGroup.GET("", h.ListProductHandler)
	ProductGroup.DELETE("/:id", h.Delete)

}

func (h *ProductHandler) ListProductHandler(c echo.Context) error {
	params := httpentity.ProductParams{}
	err := c.Bind(&params)
	if err != nil {
		return handleInvalidDataError(c, err)
	}
	res, err := h.Services.List(c.Request().Context(), params)
	return handleApplicationResponse(c, "", res, err)
}

func (h *ProductHandler) UpdateProductHandler(c echo.Context) error {

	ProductId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return handleApplicationError(c, err)
	}
	data := httpentity.UpdateProductRequest{}
	data.Id = ProductId

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

func (h ProductHandler) GetProductHandler(c echo.Context) error {

	ProductId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return handleApplicationError(c, err)
	}
	res, err := h.Services.GetByID(c.Request().Context(), ProductId)
	return handleApplicationResponse(c, "", res, err)
}

func (h ProductHandler) AddProductHandler(c echo.Context) error {
	data := httpentity.CreateProductRequest{}
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

func (h *ProductHandler) Delete(c echo.Context) error {
	addressIdStr := c.Param("id")
	addressId, err := strconv.Atoi(addressIdStr)
	if err != nil {
		return handleApplicationError(c, err)
	}

	ctx := c.Request().Context()
	err = h.Services.Delete(ctx, addressId)
	return handleApplicationResponse(c, "deleted successfully.", nil, err)
}
