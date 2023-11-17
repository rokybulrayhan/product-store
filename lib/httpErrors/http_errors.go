package httpErrors

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const (
	ErrDefault            = false
	ErrBadRequest         = "Bad request"
	ErrEmailAlreadyExists = "User with given email already exists"
	// ErrTypeAlreadyExists   = "This Type already exists"
	// ErrCommonAlreadyExists = "Already exists"
	ErrNoSuchUser       = "User not found"
	ErrWrongCredentials = "Wrong Credentials"
	ErrNotFound         = "Not Found"
	ErrUnauthorized     = "Unauthorized"
	ErrForbidden        = "Forbidden"
	ErrBadQueryParams   = "Invalid query params"
)

var (
	BadRequest            = errors.New("Bad request")
	WrongCredentials      = errors.New("Wrong Credentials")
	NotFound              = errors.New("Not Found")
	Unauthorized          = errors.New("Unauthorized")
	Forbidden             = errors.New("Forbidden")
	PermissionDenied      = errors.New("Permission Denied")
	ExpiredCSRFError      = errors.New("Expired CSRF token")
	WrongCSRFToken        = errors.New("Wrong CSRF token")
	CSRFNotPresented      = errors.New("CSRF not presented")
	NotRequiredFields     = errors.New("No such required fields")
	BadQueryParams        = errors.New("Invalid query params")
	InternalServerError   = errors.New("Internal Server Error")
	RequestTimeoutError   = errors.New("Request Timeout")
	ExistsEmailError      = errors.New("User with given email already exists")
	ExistsTypeError       = errors.New("This Type already exists")
	ExistsCommonError     = errors.New("Already exists")
	ExistsInvoiceError    = errors.New("Invoice Number Already Exists")
	ExistsProductError    = errors.New("Product with Same SKU already exits")
	ExistsNameError       = errors.New("Name already exits")
	ExistsSlugError       = errors.New("Slug already exits")
	InvalidJWTToken       = errors.New("Invalid JWT token")
	InvalidJWTClaims      = errors.New("Invalid JWT claims")
	NotAllowedImageHeader = errors.New("Not allowed image header")
	NoCookie              = errors.New("not found cookie header")
)

// Rest error interface
type RestErr interface {
	IsSuceess() bool
	Status() int
	Error() string
	Causes() interface{}
}

// Rest error struct
type RestError struct {
	ErrIsSuccess bool        `json:"success"`
	ErrStatus    int         `json:"status"`
	ErrError     string      `json:"message"`
	ErrCauses    interface{} `json:"-"`
}

// Error status
func (e RestError) IsSuceess() bool {
	return e.ErrIsSuccess
}

// Error  Error() interface method
func (e RestError) Error() string {
	return fmt.Sprintf("success: %v status: %d - errors: %s - causes: %v", e.ErrIsSuccess, e.ErrStatus, e.ErrError, e.ErrCauses)
}

// Error status
func (e RestError) Status() int {
	return e.ErrStatus
}

// RestError Causes
func (e RestError) Causes() interface{} {
	return e.ErrCauses
}

// New Rest Error
func NewRestError(status int, err string, causes interface{}) RestErr {
	return RestError{
		ErrIsSuccess: ErrDefault,
		ErrStatus:    status,
		ErrError:     err,
		ErrCauses:    causes,
	}
}

// New Rest Error
func NewCustomError(status int, err string, causes interface{}) RestErr {
	errString := ""
	errors := strings.Split(err, "custom_error!!!")
	if len(errors) >= 2 {
		errString = errors[1]
	}
	return RestError{
		ErrIsSuccess: ErrDefault,
		ErrStatus:    status,
		ErrError:     errString,
		ErrCauses:    causes,
	}
}

// New Rest Error With Message
func NewRestErrorWithMessage(status int, err string, causes interface{}) RestErr {
	return RestError{
		ErrIsSuccess: ErrDefault,
		ErrStatus:    status,
		ErrError:     err,
		ErrCauses:    causes,
	}
}

// New Rest Error From Bytes
func NewRestErrorFromBytes(bytes []byte) (RestErr, error) {
	var apiErr RestError
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}

// New Bad Request Error
func NewBadRequestError(causes interface{}) RestErr {
	return RestError{
		ErrIsSuccess: ErrDefault,
		ErrStatus:    http.StatusBadRequest,
		ErrError:     BadRequest.Error(),
		ErrCauses:    causes,
	}
}

// New Not Found Error
func NewNotFoundError(causes interface{}) RestErr {
	return RestError{
		ErrIsSuccess: ErrDefault,
		ErrStatus:    http.StatusNotFound,
		ErrError:     NotFound.Error(),
		ErrCauses:    causes,
	}
}

// New Unauthorized Error
func NewUnauthorizedError(causes interface{}) RestErr {
	return RestError{
		ErrIsSuccess: ErrDefault,
		ErrStatus:    http.StatusUnauthorized,
		ErrError:     Unauthorized.Error(),
		ErrCauses:    causes,
	}
}

// New Forbidden Error
func NewForbiddenError(causes interface{}) RestErr {
	return RestError{
		ErrIsSuccess: ErrDefault,
		ErrStatus:    http.StatusForbidden,
		ErrError:     Forbidden.Error(),
		ErrCauses:    causes,
	}
}

// New Internal Server Error
func NewInternalServerError(causes interface{}) RestErr {
	result := RestError{
		ErrIsSuccess: ErrDefault,
		ErrStatus:    http.StatusInternalServerError,
		ErrError:     InternalServerError.Error(),
		ErrCauses:    causes,
	}
	return result
}

// Parser of error string messages returns RestError
func ParseErrors(err error) RestErr {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return NewRestError(http.StatusNotFound, NotFound.Error(), err)
	case errors.Is(err, context.DeadlineExceeded):
		return NewRestError(http.StatusRequestTimeout, RequestTimeoutError.Error(), err)
	case strings.Contains(err.Error(), "SQLSTATE"):
		return parseSqlErrors(err)
	case strings.Contains(err.Error(), "Field validation") || strings.Contains(err.Error(), "schema") || strings.Contains(err.Error(), "error decoding"):
		return parseValidatorError(err)
	case strings.Contains(err.Error(), "custom_error!!!"):
		return NewCustomError(http.StatusBadRequest, err.Error(), err)
	case strings.Contains(err.Error(), "Unmarshal"):
		return NewRestError(http.StatusBadRequest, BadRequest.Error(), err)
	case strings.Contains(err.Error(), "UUID"):
		return NewRestError(http.StatusBadRequest, err.Error(), err)
	case strings.Contains(strings.ToLower(err.Error()), "cookie"):
		return NewRestError(http.StatusUnauthorized, Unauthorized.Error(), err)
	case strings.Contains(strings.ToLower(err.Error()), "token"):
		return NewRestError(http.StatusUnauthorized, Unauthorized.Error(), err)
	case strings.Contains(strings.ToLower(err.Error()), "bcrypt"):
		return NewRestError(http.StatusBadRequest, BadRequest.Error(), err)
	default:
		if restErr, ok := err.(RestErr); ok {
			return restErr
		}
		return NewInternalServerError(err)
	}
}

func parseSqlErrors(err error) RestErr {
	if strings.Contains(err.Error(), "status_name_key") {
		return NewRestError(http.StatusBadRequest, ExistsTypeError.Error(), err)
	}

	if strings.Contains(err.Error(), "products_pkey") {
		return NewRestError(http.StatusBadRequest, ExistsProductError.Error(), err)
	}

	if strings.Contains(err.Error(), "product_categories_name_key") {
		return NewRestError(http.StatusBadRequest, ExistsNameError.Error(), err)
	}

	if strings.Contains(err.Error(), "product_categories_slug_key") {
		return NewRestError(http.StatusBadRequest, ExistsSlugError.Error(), err)
	}

	if strings.Contains(err.Error(), "sales_account_invoice_number_key") {
		return NewRestError(http.StatusBadRequest, ExistsInvoiceError.Error(), err)
	}

	if strings.Contains(err.Error(), "23505") {
		return NewRestError(http.StatusBadRequest, ExistsCommonError.Error(), err)
	}

	return NewRestError(http.StatusBadRequest, BadRequest.Error(), err)
}

func parseValidatorError(err error) RestErr {
	if strings.Contains(err.Error(), "Password") {
		return NewRestError(http.StatusBadRequest, "Invalid password, min length 6", err)
	}

	if strings.Contains(err.Error(), "Email") {
		return NewRestError(http.StatusBadRequest, "Invalid email", err)
	}

	if strings.Contains(err.Error(), "invalid path") {
		return NewRestError(http.StatusBadRequest, BadQueryParams.Error(), err)
	}

	return NewRestError(http.StatusBadRequest, BadRequest.Error(), err)
}

// Error response
func ErrorResponse(err error) (int, interface{}) {
	fmt.Println("errors ------------------------->>>>>>>>>>>>>>>>>>>>>>>>>", err.Error())
	return ParseErrors(err).Status(), ParseErrors(err)
}
