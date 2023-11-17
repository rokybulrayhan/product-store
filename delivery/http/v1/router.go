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

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/uptrace/bun"
)

// Setup all routers
func SetupRouters(c *echo.Echo, conf *config.Config, db *bun.DB, jwtConfig middleware.JWTConfig, logger logger.Logger, sess *session.Session) {

	brandRepo := repository.NewBrandRepo(db)
	brandSV := brandService.NewService(brandRepo, logger)
	brandHandler := NewBrandHandler(brandSV, logger)

	v1 := c.Group("/product/store")

	health := v1.Group("/health")

	brandGroup := v1.Group("/brand")

	authenticated := middleware.JWTWithConfig(jwtConfig)
	brandHandler.MapBrandRoutes(brandGroup, authenticated)

	health.GET("", func(c echo.Context) error {
		logger.Infof("Health check RequestID: %s", lib.GetRequestID(c))
		return c.JSON(http.StatusOK, &entity.Response{
			Success: true,
			Message: "Prouct Store Server is running properly",
			Status:  http.StatusOK,
		})
	})

}
