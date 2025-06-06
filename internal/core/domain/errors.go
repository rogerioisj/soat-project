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
	InvalidOrderStatus      ErrorCode = "invalid_order_status"
	ProductNotFoundInOrder  ErrorCode = "product_not_found_in_order"
	UserAlreadyExists       ErrorCode = "user_already_exists"
	UserNotFound            ErrorCode = "user_not_found"
	InvalidOrder            ErrorCode = "invalid_order"
	InvalidItemName         ErrorCode = "invalid_item_name"
	NilItemInstance         ErrorCode = "nil_item_instance"
	InvalidItemPrice        ErrorCode = "invalid_item_price"
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
