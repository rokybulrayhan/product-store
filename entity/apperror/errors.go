package apperror

import "net/http"

var (
	LoginAuthFailed = New(http.StatusOK, "login.failed.autherror", "Email or Passwod mismtach.")

	LoginAccUnverified = New(http.StatusOK, "login.failed.unverified", "Account is not verified. Please verify account with OTP.")

	LoginAccInactive = New(http.StatusOK, "login.failed.inactive", "Account is inactive. Please contact administrator reagrding your account.")

	LoginAccLocked = New(http.StatusOK, "login.failed.locked", "Account blocked. Please contact customer support.")

	EmailExists = New(http.StatusOK, "email.duplicate", "Email already exists.")

	PhoneNumberExists = New(http.StatusOK, "phonenumber.duplicate", "Username already exists.")

	UsernameExists = New(http.StatusOK, "username.duplicate", "Username already exists")

	GenericNotFound = New(http.StatusNotFound, "record.not.found", "Record not found")

	UserNotFound = New(http.StatusNotFound, "user.not.found", "User not found")

	SessionNotFound = New(http.StatusNotFound, "session.not.found", "Session not found")

	AddressNotFound = New(http.StatusNotFound, "address.not.found", "Address not found")

	WrongCode = New(http.StatusOK, "verification.failed", "Verfication code did not match")

	NoCode = New(http.StatusNotFound, "verfication.not.found", "No such code found")

	TokenInvalid = New(http.StatusBadRequest, "token.invalid", "Invalid refresh token provided.")

	CodeExpired = New(http.StatusOK, "verfication.code.expired", "Your OTP expired. Please go back and try logging in again")

	PasswordMismatch = New(http.StatusOK, "password.error", "Old password did not match. Please correct your old password")
	// New("","")

	UsernamePasswordMismatch = New(http.StatusOK, "login.failed.autherror", "Phone number or Passwod mismtach.")
	CantSendOtp              = New(http.StatusOK, "otp.failed", "Unable to send OTP at this moment. Please try again after a minute.")

	InteralError = New(http.StatusInternalServerError, "error.internal", "Internal error occured. Please contact webmaster.")
	Unknown      = New(http.StatusInternalServerError, "error.unknown", "Something went wrong. Please try again after some time. Contact customer support for more information.")
)

func New(httpCode int, errorCode string, errorMessage string) *ApplicationError {
	return &ApplicationError{
		HTTPCode:     httpCode,
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
	}
}

type ApplicationError struct {
	HTTPCode     int    `json:"http_code"`
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	ActualError  error  `json:"-"`
}

func (a *ApplicationError) Error() string {
	return a.ErrorCode
}

func (a *ApplicationError) Wrap(err error) *ApplicationError {
	a.ActualError = err
	return a
}
