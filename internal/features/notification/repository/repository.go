package notification_repository

import "net/http"

type NotificationRepository struct {
	token  string
	client *http.Client
}

func NewNotificationRepository(token string) *NotificationRepository {
	return &NotificationRepository{
		token:  token,
		client: &http.Client{},
	}
}
