package er

import (
	"strconv"

	"github.com/google/uuid"
)

// Used on internal error
type AppError struct {
	code    code
	traceID uuid.UUID // used to trace log
}

// Get error message from code, if not found, return ""
func (e *AppError) Error() string {
	msg, ok := CodeToMsg[e.code]
	if !ok {
		return ""
	}
	return msg
}

func (e *AppError) Code() code {
	return e.code
}

func (e *AppError) HTTPCode() int {
	codeStr, _ := strconv.Atoi(e.code.String()[:3])
	return codeStr
}

// NewAppError will log the caller stack and original error message if exist
func NewAppErr(code code) error {
	return &AppError{
		code:    code,
		traceID: logAppErr(code),
	}
}
