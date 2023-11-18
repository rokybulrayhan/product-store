package v1

import (
	"fmt"
	"strconv"

	"github.com/techno/entity/httpentity"
	"github.com/techno/lib/logger"
	Brand "github.com/techno/service/brand"

	"github.com/labstack/echo/v4"
)

type BrandHandler struct {
	Services *Brand.Service
	Logger   logger.Logger
}

func NewBrandHandler(services *Brand.Service, logger logger.Logger) *BrandHandler {
	return &BrandHandler{
		Services: services,
		Logger:   logger,
	}
}

// Map Brands routes
func (h BrandHandler) MapBrandRoutes(BrandGroup *echo.Group, authenticated echo.MiddlewareFunc) {
	BrandGroup.POST("", h.AddBrandHandler)
	BrandGroup.PUT("/:id", h.UpdateBrandHandler)
	BrandGroup.GET("/:id", h.GetBrandHandler)
	BrandGroup.GET("", h.ListBrandHandler)
	BrandGroup.DELETE("/:id", h.Delete)

}

func (h *BrandHandler) ListBrandHandler(c echo.Context) error {
	params := httpentity.BrandParams{}
	err := c.Bind(&params)
	if err != nil {
		return handleInvalidDataError(c, err)
	}
	res, err := h.Services.List(c.Request().Context(), params)
	return handleApplicationResponse(c, "", res, err)
}

func (h *BrandHandler) UpdateBrandHandler(c echo.Context) error {

	BrandId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return handleApplicationError(c, err)
	}
	data := httpentity.UpdateBrandRequest{}
	data.Id = BrandId

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

func (h BrandHandler) GetBrandHandler(c echo.Context) error {

	BrandId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return handleApplicationError(c, err)
	}
	res, err := h.Services.GetByID(c.Request().Context(), BrandId)
	return handleApplicationResponse(c, "", res, err)
}

func (h BrandHandler) AddBrandHandler(c echo.Context) error {
	data := httpentity.CreateBrandRequest{}
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

func (h *BrandHandler) Delete(c echo.Context) error {
	addressIdStr := c.Param("id")
	addressId, err := strconv.Atoi(addressIdStr)
	if err != nil {
		return handleApplicationError(c, err)
	}

	ctx := c.Request().Context()
	err = h.Services.Delete(ctx, addressId)
	return handleApplicationResponse(c, "deleted successfully.", nil, err)
}
