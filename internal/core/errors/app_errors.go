package errors

type AppError struct {
	Code    string
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

func NotFound(msg string) *AppError {
	return &AppError{
		Code:    "NOT_FOUND",
		Message: msg,
	}
}

func Duplicate(msg string) *AppError {
	return &AppError{
		Code:    "DUPLICATE",
		Message: msg,
	}
}

func Invalid(msg string) *AppError {
	return &AppError{
		Code:    "INVALID",
		Message: msg,
	}
}

func Unauthorized(msg string) *AppError {
	return &AppError{
		Code:    "UNAUTHORIZED",
		Message: msg,
	}
}
