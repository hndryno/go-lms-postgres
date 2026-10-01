package response

type Response struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func Success(code int, message string, data any) Response {
	return Response{
		Code:    code,
		Success: true,
		Message: message,
		Data:    data,
	}
}

func Error(code int, message string) Response {
	return Response{
		Code:    code,
		Success: false,
		Message: message,
	}
}
