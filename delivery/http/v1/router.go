package v1

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-contact-service/config"
	"github.com/go-contact-service/entity"
	"github.com/go-contact-service/lib"
	"github.com/go-contact-service/lib/logger"
	"github.com/go-contact-service/repository"

	brandService "github.com/go-contact-service/service/brand"
	categoryService "github.com/go-contact-service/service/category"
	productService "github.com/go-contact-service/service/product"
	productStockService "github.com/go-contact-service/service/productstock"
	supplierService "github.com/go-contact-service/service/supplier"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/uptrace/bun"
)

// Setup all routers
func SetupRouters(c *echo.Echo, conf *config.Config, db *bun.DB, jwtConfig middleware.JWTConfig, logger logger.Logger, sess *session.Session) {

	brandRepo := repository.NewBrandRepo(db)
	brandSV := brandService.NewService(brandRepo, logger)
	brandHandler := NewBrandHandler(brandSV, logger)

	productRepo := repository.NewProductRepo(db)
	productSV := productService.NewService(productRepo, logger)
	productHandler := NewProductHandler(productSV, logger)

	productStockRepo := repository.NewProductStockRepo(db)
	productStockSV := productStockService.NewService(productStockRepo, logger)
	productStockHandler := NewProductStockHandler(productStockSV, logger)

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
