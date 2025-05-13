package er

// Use string as code to support more flexibility than float
type code string

func (c code) String() string {
	return string(c)
}

// The int part should follow http status code
const (
	BadRequest      code = "400.0000"
	EmptyRequest    code = "400.0001"
	ParsePayload    code = "400.0002"
	ValidateInput   code = "400.0003"
	AlreadyResetOTP code = "400.0004"
	AlreadyExists   code = "400.0005"
	InvalidPassword code = "400.0006"
	InvalidEmail    code = "400.0007"
	InvalidCode     code = "400.0009"

	Unauthorized               code = "401.0000"
	NewPasswordRequired        code = "401.0001"
	MissingAuthorizationHeader code = "401.0002"

	NotFound code = "404.0000"

	Unknown       code = "500.0000"
	DBUnavailable code = "500.0001"

	NotImplemented code = "501.0000"
)

var CodeToMsg = map[code]string{
	BadRequest:      "bad request",
	EmptyRequest:    "empty request body",
	ParsePayload:    "parse payload error",
	ValidateInput:   "validate input error",
	AlreadyResetOTP: "already reset otp",
	AlreadyExists:   "already exists",
	InvalidPassword: "invalid password",
	InvalidEmail:    "invalid email",
	InvalidCode:     "invalid code",

	Unauthorized:               "unauthorized",
	NewPasswordRequired:        "new password required",
	MissingAuthorizationHeader: "missing authorization header",

	NotFound: "not found",

	Unknown:       "unknown",
	DBUnavailable: "database unavailable",

	NotImplemented: "not implemented",
}
