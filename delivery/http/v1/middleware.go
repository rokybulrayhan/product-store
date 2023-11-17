package v1

import (
	"log"
	"net/http"

	"github.com/go-contact-service/entity"
	"github.com/go-contact-service/entity/httpentity"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slices"
)

func InvalidJwt(err error) error {
	if err != nil {
		log.Println("Got from echo")
	}
	return echo.NewHTTPError(http.StatusUnauthorized,
		&httpentity.ErrorResponse{
			Success:      false,
			ErrorCode:    "unauthorized",
			ErrorMessage: "You are not authorized to access this resource",
		})
}

func authorize(allowCustomers bool, internalCategories ...int) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claim, ok := c.Get("user").(*jwt.Token).Claims.(*entity.JwtClaim)
			//claim := claims.Claims.(*entity.JwtClaim)
			if !ok {
				return InvalidJwt(nil)
			}
			if !allowCustomers && claim.AccountType == entity.ACCOUNT_TYPE_CUSTOMER {
				return InvalidJwt(nil)
			}
			if claim.AccountType == entity.ACCOUNT_TYPE_INTERNAL && len(internalCategories) != 0 && !slices.Contains(internalCategories, claim.AccountCategoryId) {
				return InvalidJwt(nil)
			}
			c.Set("id", claim.CustomClaim.Id)
			c.Set("account_category_id", claim.AccountCategoryId)
			c.Set("account_type", claim.AccountType)
			return next(c)
		}
	}
}
