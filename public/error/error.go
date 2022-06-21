package error

import "strconv"

// type internalError struct {
// 	*Err
// }

func New(code int, message string, domain string, reason string, suggestion string) Error {
	return Error{
		Code:       code,
		Message:    message,
		Domain:     domain,
		Reason:     reason,
		Suggestion: suggestion,
	}
}

func (e *Error) Error() string {
	return e.Domain + ";" + e.Message + ":" + strconv.Itoa(e.Code) + ";" + e.Reason + ";" + e.Suggestion
}

func NewDomain(domain string) CreateDomainErrorFunc {
	return func(code int, message, reason, suggestion string) Error {
		return New(code, message, domain, reason, suggestion)
	}
}
