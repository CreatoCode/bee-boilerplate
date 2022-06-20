package error

type Error struct {
	Code       int    `json:"code"`       // code is the error
	Message    string `json:"message"`    // message for user
	Domain     string `json:"domain"`     // described by names that are arbitrary strings used to differentiate groups of codes;
	Reason     string `json:"reason"`     // describing why the operation failed
	Suggestion string `json:"suggestion"` // describing what the user can do to fix the problem
}

type IError interface {
	Error() string
}
