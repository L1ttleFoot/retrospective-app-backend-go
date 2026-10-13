package notification_repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type SendNotificationRequest struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

type SendNotificationResponse struct {
	OK          bool   `json:"ok"`
	Description string `json:"description"`
}

func (r *NotificationRepository) SendNotification(text string) error {

	requestBody := SendNotificationRequest{
		ChatID: 368161962,
		Text:   text,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("marshal request body: %w", err)
	}

	url := fmt.Sprintf(
		"https://api.telegram.org/bot%s/sendMessage",
		r.token,
	)

	ctx := context.Background()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := r.client.Do(req)
	if err != nil {
		return fmt.Errorf("send telegram request: %w", err)
	}
	defer resp.Body.Close()

	var notificationResponse SendNotificationResponse

	if err := json.NewDecoder(resp.Body).Decode(&notificationResponse); err != nil {
		return fmt.Errorf("decode telegram response: %w", err)
	}

	fmt.Println(notificationResponse, resp.StatusCode != http.StatusOK || !notificationResponse.OK)

	if resp.StatusCode != http.StatusOK || !notificationResponse.OK {
		return fmt.Errorf(
			"telegram API error: status=%d description=%s",
			resp.StatusCode,
			notificationResponse.Description,
		)
	}

	return nil
}
