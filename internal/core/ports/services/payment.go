package services

type PaymentServiceInterface interface {
	GenerateQRCode(value int64) error
}
