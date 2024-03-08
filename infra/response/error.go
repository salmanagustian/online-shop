package response

import (
	"errors"
	"net/http"
)

// error general
var (
	ErrNotFound        = errors.New("not found")
	ErrUnauthorized    = errors.New("Unauthorized")
	ErrForbiddenAccess = errors.New("forbiden access")
)

var (
	// auth
	ErrEmailRequired         = errors.New("email is required")
	ErrEmailInvalid          = errors.New("email is invalid")
	ErrPasswordRequired      = errors.New("password is required")
	ErrPasswordInvalidLength = errors.New("password must have mininum 6 characters")
	ErrAuthIsNotExists       = errors.New("auth is not exists")
	ErrEmailAlreadyUsed      = errors.New("email already used")
	ErrPasswordNotMatch      = errors.New("password not match")

	// products
	ErrProductRequired = errors.New("product is required")
	ErrProductInvalid  = errors.New("product must have minimum 4 characters")
	ErrStockInvalid    = errors.New("stock must be greater than 0")
	ErrPriceInvalid    = errors.New("price must be greater than 0")

	// transactions
	ErrAmountInvalid         = errors.New("invalid amount")
	ErrAmoutGreaterThanStock = errors.New("amount greater than stock")
)

type Error struct {
	Message  string
	Code     string
	HttpCode int
}

func NewError(msg string, code string, httpCode int) Error {
	return Error{
		Message:  msg,
		Code:     code,
		HttpCode: httpCode,
	}
}

func (e Error) Error() string {
	return e.Message
}

// error general
var (
	ErrorGeneral         = NewError("general error", "99999", http.StatusInternalServerError)
	ErrorBadRequest      = NewError("bad request", "40000", http.StatusBadRequest)
	ErrorNotFound        = NewError(ErrNotFound.Error(), "40400", http.StatusNotFound)
	ErrorForbiddenAccess = NewError(ErrForbiddenAccess.Error(), "40100", http.StatusForbidden)
)

var (
	// error bad request
	ErrorEmailRequired         = NewError(ErrEmailRequired.Error(), "40001", http.StatusBadRequest)
	ErrorEmailInvalid          = NewError(ErrEmailInvalid.Error(), "40002", http.StatusBadRequest)
	ErrorPasswordRequired      = NewError(ErrPasswordRequired.Error(), "40003", http.StatusBadRequest)
	ErrorPasswordInvalidLength = NewError(ErrPasswordInvalidLength.Error(), "40004", http.StatusBadRequest)
	ErrorProductRequired       = NewError(ErrProductRequired.Error(), "40005", http.StatusBadRequest)
	ErrorProductInvalid        = NewError(ErrProductInvalid.Error(), "40006", http.StatusBadRequest)
	ErrorStockInvalid          = NewError(ErrStockInvalid.Error(), "40007", http.StatusBadRequest)
	ErrorPriceInvalid          = NewError(ErrPriceInvalid.Error(), "40008", http.StatusBadRequest)
	ErrorInvalidAmount         = NewError(ErrAmountInvalid.Error(), "40009", http.StatusBadRequest)

	ErrorAuthIsNotExist   = NewError(ErrAuthIsNotExists.Error(), "40401", http.StatusNotFound)
	ErrorEmailAlreadyUsed = NewError(ErrEmailAlreadyUsed.Error(), "40901", http.StatusConflict)
	ErrorUnauthorized     = NewError(ErrUnauthorized.Error(), "40100", http.StatusUnauthorized)
	ErrorPasswordNotMatch = NewError(ErrPasswordNotMatch.Error(), "40101", http.StatusUnauthorized)
)

var (
	ErrorMapping = map[string]Error{
		ErrNotFound.Error():              ErrorNotFound,
		ErrEmailRequired.Error():         ErrorEmailRequired,
		ErrEmailInvalid.Error():          ErrorEmailInvalid,
		ErrPasswordRequired.Error():      ErrorPasswordRequired,
		ErrPasswordInvalidLength.Error(): ErrorPasswordInvalidLength,
		ErrAuthIsNotExists.Error():       ErrorAuthIsNotExist,
		ErrEmailAlreadyUsed.Error():      ErrorEmailAlreadyUsed,
		ErrorUnauthorized.Error():        ErrorUnauthorized,
		ErrPasswordNotMatch.Error():      ErrorPasswordNotMatch,
		ErrForbiddenAccess.Error():       ErrorForbiddenAccess,
	}
)
