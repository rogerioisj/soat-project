package domain

type ErrorCode string

const (
	InvalidNameRange   ErrorCode = "invalid_name_range"
	InvalidEmailRange  ErrorCode = "invalid_email_range"
	InvalidEmailFormat ErrorCode = "invalid_email_format"
	InvalidCPF         ErrorCode = "invalid_cpf"
	InvalidUserID      ErrorCode = "invalid_user_id"
)

type DomainError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

func NewDomainError(code ErrorCode, message string) *DomainError {
	return &DomainError{
		Code:    code,
		Message: message,
	}
}

func (e *DomainError) Error() string {
	return e.Message
}

func (e *DomainError) GetCode() ErrorCode {
	return e.Code
}

func (e *DomainError) GetMessage() string {
	return e.Message
}

func (e *DomainError) Is(code ErrorCode) bool {
	return e.Code == code
}
