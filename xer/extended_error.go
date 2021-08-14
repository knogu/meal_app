package xer

// extended error

type ErrType struct {
	StatusCode int
	Summary    string
}

var (
	JsonFormatInvalid = ErrType{StatusCode: 400, Summary: "request json format is invalid. Failed to parse json"}
	ParamInvalid      = ErrType{StatusCode: 400, Summary: "request json validation failed"}
	WrongPassword     = ErrType{StatusCode: 400, Summary: "password is wrong"}

	TeamNotFound  = ErrType{StatusCode: 404, Summary: "team not found"}
	EventNotFound = ErrType{StatusCode: 404, Summary: "event not found"}
	UserNotFound  = ErrType{StatusCode: 404, Summary: "user not found"}

	NotAuthorized    = ErrType{StatusCode: 403, Summary: "you are not authorized"}
	MethodNotAllowed = ErrType{StatusCode: 405, Summary: "Method not allowed"}
)

type Err4xx struct {
	ErrType
	Detail string
}

func (err Err4xx) Error() string {
	return "Summary: " + err.Summary + " Detail: " + err.Detail
}
