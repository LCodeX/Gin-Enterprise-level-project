package resp

type StatusCode struct {
	Code    int
	Message string
}

var (
	BadRequest          = StatusCode{Code: 400, Message: "Bad Request"}
	Unauthorized        = StatusCode{Code: 401, Message: "Unauthorized"}
	Forbidden           = StatusCode{Code: 403, Message: "Forbidden"}
	NotFound            = StatusCode{Code: 404, Message: "Not Found"}
	InternalServerError = StatusCode{Code: 500, Message: "Internal Server Error"}
	ServiceUnavailable  = StatusCode{Code: 503, Message: "Service Unavailable"}

	Success            = StatusCode{Code: 1000, Message: "Ok"}
	UserAlreadyExists  = StatusCode{Code: 1001, Message: "User already exists"}
	UserNotFound       = StatusCode{Code: 1002, Message: "User not found"}
	InvalidCredentials = StatusCode{Code: 1003, Message: "Invalid username or password"}
	TokenExpired       = StatusCode{Code: 1004, Message: "Token has expired"}
	RequestDataError   = StatusCode{Code: 1005, Message: "Request data error"}
)

func (s StatusCode) GetMessage() string {
	return s.Message
}

func (s StatusCode) GetCode() int {
	return s.Code
}
