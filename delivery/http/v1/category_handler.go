package v1

import (
	"fmt"
	"strconv"

	"github.com/techno/entity/httpentity"
	"github.com/techno/lib/logger"
	Category "github.com/techno/service/category"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	Services *Category.Service
	Logger   logger.Logger
}

func NewCategoryHandler(services *Category.Service, logger logger.Logger) *CategoryHandler {
	return &CategoryHandler{
		Services: services,
		Logger:   logger,
	}
}

// Map Categorys routes
func (h CategoryHandler) MapCategoryRoutes(CategoryGroup *echo.Group, authenticated echo.MiddlewareFunc) {
	CategoryGroup.POST("", h.AddCategoryHandler)
	CategoryGroup.PUT("/:id", h.UpdateCategoryHandler)
	CategoryGroup.GET("/:id", h.GetCategoryHandler)
	CategoryGroup.GET("", h.ListCategoryHandler)
	CategoryGroup.DELETE("/:id", h.Delete)

}

func (h *CategoryHandler) ListCategoryHandler(c echo.Context) error {
	params := httpentity.CategoryParams{}
	err := c.Bind(&params)
	if err != nil {
		return handleInvalidDataError(c, err)
	}
	res, err := h.Services.List(c.Request().Context(), params)
	return handleApplicationResponse(c, "", res, err)
}

func (h *CategoryHandler) UpdateCategoryHandler(c echo.Context) error {

	CategoryId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return handleApplicationError(c, err)
	}
	data := httpentity.UpdateCategoryRequest{}
	data.Id = CategoryId

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

func (h CategoryHandler) GetCategoryHandler(c echo.Context) error {

	CategoryId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return handleApplicationError(c, err)
	}
	res, err := h.Services.GetByID(c.Request().Context(), CategoryId)
	return handleApplicationResponse(c, "", res, err)
}

func (h CategoryHandler) AddCategoryHandler(c echo.Context) error {
	data := httpentity.CreateCategoryRequest{}
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

func (h *CategoryHandler) Delete(c echo.Context) error {
	addressIdStr := c.Param("id")
	addressId, err := strconv.Atoi(addressIdStr)
	if err != nil {
		return handleApplicationError(c, err)
	}

	ctx := c.Request().Context()
	err = h.Services.Delete(ctx, addressId)
	return handleApplicationResponse(c, "deleted successfully.", nil, err)
}
