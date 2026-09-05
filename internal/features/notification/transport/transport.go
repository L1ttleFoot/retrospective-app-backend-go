package notification_transport

import "golang-clean-server/internal/core/domain"

type NotificationHandler struct {
	telegramService TelegramService
}

type TelegramService interface {
	SendNotification(notification domain.Notification) error
}

func NewTelegramHandler(telegramService TelegramService) *NotificationHandler {
	return &NotificationHandler{telegramService: telegramService}
}
