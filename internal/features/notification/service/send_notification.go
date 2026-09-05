package notification_service

import (
	"fmt"
	"golang-clean-server/internal/core/domain"
)

func (h *NotificationService) SendNotification(notification domain.Notification) error {

	text := fmt.Sprintf(
		"Новое сообщение:\n%s\n\nСоздано: %s",
		notification.Text,
		notification.CreatedAt.Format("02.01.2006 15:04"),
	)

	if err := h.notificationRepository.SendNotification(text); err != nil {
		return fmt.Errorf("send notification: %w", err)
	}

	return nil
}
