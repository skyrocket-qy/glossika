package er

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

// Wrap error to AppError
func W(err error, codes ...code) error {
	if err == nil {
		return nil
	}

	var appErr *AppError
	if errors.As(err, &appErr) {
		// already an AppError, do nothing
		return err
	}

	var code code = Unknown
	if len(codes) > 0 {
		code = codes[0]
	}

	return &AppError{
		code:    code,
		traceID: logAppErr(code, err),
	}
}

func logAppErr(code code, errs ...error) (traceID uuid.UUID) {
	callStk := getCallStack(3)

	traceID = uuid.New()

	event := log.Error().
		Str("traceID", traceID.String()).
		Str("code", code.String())

	if len(errs) > 0 {
		event = event.Err(errs[0])
	}

	event.Msg(callStk)

	return traceID
}

func getCallStack(callerSkip ...int) (stkMsg string) {
	pc := make([]uintptr, 10)
	skip := 2
	if len(callerSkip) > 0 {
		skip = callerSkip[0]
	}
	n := runtime.Callers(skip, pc)

	frames := runtime.CallersFrames(pc[:n-2])

	stkMsg = "Call stack:\n"

	for {
		frame, more := frames.Next()
		stkMsg += fmt.Sprintf("%s\n\t%s:%d\n", frame.Function, frame.File, frame.Line)
		if !more {
			break
		}
	}

	return stkMsg
}
