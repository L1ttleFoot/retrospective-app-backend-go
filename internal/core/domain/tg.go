package domain

import "time"

type Notification struct {
	Text      string
	CreatedAt time.Time
}
