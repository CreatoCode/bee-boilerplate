package error

// 错误需要的字段：错误码，错误消息，所属域，错误原因，错误解决建议
type Error struct {
	Code       int    `json:"code"`       // code is the error
	Message    string `json:"message"`    // message of the error
	Domain     string `json:"domain"`     // domain of the error
	Reason     string `json:"reason"`     // what is the error
	Suggestion string `json:"suggestion"` // fix suggestion
}

type IError interface {
	Error() string
}
