package services

type NotificationServiceInterface interface {
	SendNotification(message string) error
}
