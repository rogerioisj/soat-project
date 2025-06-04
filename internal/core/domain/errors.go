package domain

type ErrorCode string

const (
	InvalidNameRange        ErrorCode = "invalid_name_range"
	InvalidEmailRange       ErrorCode = "invalid_email_range"
	InvalidEmailFormat      ErrorCode = "invalid_email_format"
	InvalidCPF              ErrorCode = "invalid_cpf"
	InvalidDescriptionRange ErrorCode = "invalid_description_range"
	InvalidPriceRange       ErrorCode = "invalid_price_range"
	InvalidProductType      ErrorCode = "invalid_product_type"
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

func (e *DomainError) Is(code ErrorCode) bool {
	return e.Code == code
}
