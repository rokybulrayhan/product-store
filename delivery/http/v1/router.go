package v1

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/techno/config"
	"github.com/techno/entity"
	"github.com/techno/lib"
	"github.com/techno/lib/logger"
	"github.com/techno/repository"

	brandService "github.com/techno/service/brand"
	categoryService "github.com/techno/service/category"
	productService "github.com/techno/service/product"
	productStockService "github.com/techno/service/productstock"
	supplierService "github.com/techno/service/supplier"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/uptrace/bun"
)

// Setup all routers
func SetupRouters(c *echo.Echo, conf *config.Config, db *bun.DB, jwtConfig middleware.JWTConfig, logger logger.Logger, sess *session.Session) {

	brandRepo := repository.NewBrandRepo(db)
	brandSV := brandService.NewService(brandRepo, logger)
	brandHandler := NewBrandHandler(brandSV, logger)

	productStockRepo := repository.NewProductStockRepo(db)
	productStockSV := productStockService.NewService(productStockRepo, logger)
	productStockHandler := NewProductStockHandler(productStockSV, logger)

	productRepo := repository.NewProductRepo(db)
	productSV := productService.NewService(productRepo, logger, *productStockSV)
	productHandler := NewProductHandler(productSV, logger)

	categoryRepo := repository.NewCategoryRepo(db)
	categorySV := categoryService.NewService(categoryRepo, logger)
	categoryHandler := NewCategoryHandler(categorySV, logger)

	supplierRepo := repository.NewSupplierRepo(db)
	supplierSV := supplierService.NewService(supplierRepo, logger)
	supplierHandler := NewSupplierHandler(supplierSV, logger)

	v1 := c.Group("/product/store")

	health := v1.Group("/health")

	brandGroup := v1.Group("/brand")
	productStockGroup := v1.Group("/product/stock")
	productGroup := v1.Group("/product")
	supplierGroup := v1.Group("/supplier")
	categoryGroup := v1.Group("/category")

	authenticated := middleware.JWTWithConfig(jwtConfig)

	brandHandler.MapBrandRoutes(brandGroup, authenticated)
	productHandler.MapProductRoutes(productGroup, authenticated)
	productStockHandler.MapProductStockRoutes(productStockGroup, authenticated)
	supplierHandler.MapSupplierRoutes(supplierGroup, authenticated)
	categoryHandler.MapCategoryRoutes(categoryGroup, authenticated)

	health.GET("", func(c echo.Context) error {
		logger.Infof("Health check RequestID: %s", lib.GetRequestID(c))
		return c.JSON(http.StatusOK, &entity.Response{
			Success: true,
			Message: "Prouct Store Server is running properly",
			Status:  http.StatusOK,
		})
	})

}
