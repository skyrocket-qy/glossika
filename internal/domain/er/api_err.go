package er

// Used on http response
type APIError struct {
	TraceID string `json:"traceID"`
	Code    string `json:"code"`
	Msg     string `json:"msg"`
}

func NewAPIErr(traceID string, code string, msg string) *APIError {
	return &APIError{
		TraceID: traceID,
		Code:    code,
		Msg:     msg,
	}
}
