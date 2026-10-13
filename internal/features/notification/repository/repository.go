package notification_repository

import (
	core_telegram "golang-clean-server/internal/core/telegram"
	"net/http"
	"time"
)

type NotificationRepository struct {
	token  string
	client *http.Client
}

func NewNotificationRepository(cfg core_telegram.Config) *NotificationRepository {
	return &NotificationRepository{
		token:  cfg.Token,
		client: &http.Client{Timeout: 10 * time.Second},
	}
}
