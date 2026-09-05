package notification_transport

import (
	"encoding/json"
	"fmt"
	"golang-clean-server/internal/core/domain"
	"time"
)

type MessageCreatedEvent struct {
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}

func (h *NotificationHandler) SendNotification(body []byte) error {

	var event MessageCreatedEvent

	if err := json.Unmarshal(body, &event); err != nil {
		return fmt.Errorf("unmarshal telegram message: %w", err)
	}

	notification := domain.Notification{
		Text:      event.Text,
		CreatedAt: event.CreatedAt,
	}

	if err := h.telegramService.SendNotification(notification); err != nil {
		return fmt.Errorf("send notification: %w", err)
	}

	return nil
}
