package notification_service

type NotificationService struct {
	notificationRepository NotificationRepository
}

type NotificationRepository interface {
	SendNotification(text string) error
}

func NewNotificationService(notificationRepository NotificationRepository) *NotificationService {
	return &NotificationService{notificationRepository: notificationRepository}
}
