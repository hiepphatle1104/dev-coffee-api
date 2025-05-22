package common

type Response struct {
	Data    any          `json:"data"`
	Message string       `json:"message"`
	Success bool         `json:"success"`
	Error   *CustomError `json:"error"`
	Paging  *Paging      `json:"paging,omitempty"`
}

func NewResponse(data any, message string, success bool, err *CustomError) *Response {
	return &Response{
		Data:    data,
		Message: message,
		Success: success,
		Error:   err,
	}
}

func NewSuccessResponse(data any) *Response {
	return &Response{
		Data:    data,
		Message: "Request success.",
		Success: true,
	}
}

func NewSuccessResponseWithPaging(data any, paging *Paging) *Response {
	return &Response{
		Data:    data,
		Message: "Request success.",
		Success: true,
		Paging:  paging,
	}
}

func NewSuccessCreatedResponse(data any) *Response {
	return &Response{
		Data:    data,
		Message: "Request success.",
		Success: true,
	}
}

func NewDatabaseErrorResponse(err error) *Response {
	return &Response{
		Data:    nil,
		Error:   NewCustomError(err.Error(), DB_CONN_ERROR),
		Message: "Request failed.",
		Success: false,
	}
}

func NewBadRequestErrorResponse(err error) *Response {
	return &Response{
		Data:    nil,
		Error:   NewCustomError(err.Error(), BAD_REQUEST),
		Message: "Request failed.",
		Success: false,
	}
}

func NewErrorResponse(err error) *Response {
	return &Response{
		Data:    nil,
		Error:   NewCustomError(err.Error(), INTERNAL_SERVER_ERROR),
		Message: "Request failed.",
		Success: false,
	}
}
