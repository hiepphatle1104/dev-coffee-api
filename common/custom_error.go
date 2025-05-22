package common

type CustomError struct {
	Message string     `json:"message"`
	Code    CustomCode `json:"code"`
}

func NewCustomError(message string, code CustomCode) *CustomError {
	return &CustomError{
		Message: message,
		Code:    code,
	}
}

type CustomCode string

const (
	INTERNAL_SERVER_ERROR CustomCode = "INTERNAL_SERVER_ERROR"
	DB_CONN_ERROR         CustomCode = "DB_CONN_ERROR"
	ITEM_NOT_FOUND        CustomCode = "ITEM_NOT_FOUND"
	ORDER_NOT_FOUND       CustomCode = "ORDER_NOT_FOUND"
	PAYMENT_NOT_FOUND     CustomCode = "PAYMENT_NOT_FOUND"
	BAD_REQUEST           CustomCode = "BAD_REQUEST"
)
