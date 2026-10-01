package response

type Response struct {
	Code       int         `json:"code"`
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       any         `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

func Success(code int, message string, data any) Response {
	return Response{
		Code:    code,
		Success: true,
		Message: message,
		Data:    data,
	}
}

func SuccessWithPagination(
	code int,
	message string,
	data any,
	pagination Pagination,
) Response {
	return Response{
		Code:       code,
		Success:    true,
		Message:    message,
		Data:       data,
		Pagination: &pagination,
	}
}

func Error(code int, message string) Response {
	return Response{
		Code:    code,
		Success: false,
		Message: message,
	}
}
