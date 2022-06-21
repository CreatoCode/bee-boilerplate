package error

const (
	CodeInvalidArgument = iota + 1 // pass an invalid argument to a method
)

type Error struct {
	Code       int    `json:"code"`       // code is the error
	Message    string `json:"msg"`        // message for user
	Domain     string `json:"domain"`     // described by names that are arbitrary strings used to differentiate groups of codes;
	Reason     string `json:"reason"`     // describing why the operation failed
	Suggestion string `json:"suggestion"` // describing what the user can do to fix the problem
}

type CreateDomainErrorFunc = func(code int, message string, reason string, suggestion string) Error

type IErr interface {
	Error() string
	NewDomain(domain string) CreateDomainErrorFunc
	New(code int, message string, domain string, reason string, suggestion string) Error
}
