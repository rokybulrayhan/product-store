package v1

import (
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/mailru/easyjson"
	"github.com/techno/entity"
	"github.com/techno/entity/apperror"
	"github.com/techno/entity/httpentity"
)

func sessionData(c echo.Context) *entity.JwtClaim {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*entity.JwtClaim)
	return claims
}

func GetCustomerIdFromToken(c echo.Context) string {
	// If internal user is browsing, return customer ID from query param
	if c.Get("account_type").(string) == entity.ACCOUNT_TYPE_INTERNAL {
		return c.QueryParam("id")
	}
	return c.Get("id").(string)
}
func GetInternlUserIdFromTokem(c echo.Context) string {
	return c.Get("id").(string)
}

func GetAccountTypeFromToken(c echo.Context) string {
	return c.Get("account_type").(string)
}

// func handleError(c echo.Context, err ...error) error {
// 	c.JSON(http.StatusInternalServerError, err)
// 	return nil
// }

func handleValidationError(c echo.Context, validationErrors []httpentity.FieldError) error {
	return c.JSON(http.StatusBadRequest, &httpentity.ErrorResponse{
		Success:      false,
		ErrorCode:    "field.validation",
		ErrorMessage: "Field validation error",
		FieldErrors:  validationErrors,
	})
}

func handleInvalidDataError(c echo.Context, err error) error {
	// TODO: Log error
	return c.JSON(http.StatusBadRequest, &httpentity.ErrorResponse{
		Success:      false,
		ErrorCode:    "payload.invalid",
		ErrorMessage: "Invalid data provided",
	})
}

func handleApplicationError(c echo.Context, err error) error {

	appError, ok := err.(*apperror.ApplicationError)
	if ok {
		// Log original error here if needed
		if appError.ActualError != nil {
			log.Println(appError.ActualError.Error())
		}

		json, _ := easyjson.Marshal(httpentity.ErrorResponse{
			Success:      false,
			ErrorCode:    appError.ErrorCode,
			ErrorMessage: appError.ErrorMessage,
		})
		return c.Blob(appError.HTTPCode, "application/json", json)
		// return c.JSON(http.StatusOK, &httpentity.ErrorResponse{
		// 	Success:      false,
		// 	ErrorCode:    appError.ErrorCode,
		// 	ErrorMessage: appError.ErrorMessage,
		// })
	}
	log.Println(err.Error())
	json, _ := easyjson.Marshal(httpentity.ErrorResponse{

		ErrorCode:    "error.unknown",
		ErrorMessage: "Something went wrong. Please try again after some time. Contact customer support for more information.",
		Success:      false,
	})
	return c.Blob(http.StatusInternalServerError, "application/json", json)
	// return c.JSON(http.StatusInternalServerError, httpentity.ErrorResponse{
	// 	Success:      false,
	// 	ErrorCode:    "error.unknown",
	// 	ErrorMessage: "Something went wrong. Please try again after some time. Contact customer support for more information.",
	// })

}

func handleApplicationResponse(c echo.Context, successMessage string, res interface{}, err error) error {
	if err != nil {
		return handleApplicationError(c, err)
	}
	json, _ := easyjson.Marshal(httpentity.OkResponse{
		Success: true,
		Message: successMessage,
		Data:    res,
	})
	return c.Blob(http.StatusOK, "application/json", json)
	// return c.JSON(http.StatusOK, httpentity.OkResponse{
	// 	Success: true,
	// 	Message: successMessage,
	// 	Data:    res,
	// })

}
