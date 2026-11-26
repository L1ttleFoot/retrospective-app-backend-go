package exp_gateway

import (
	"net/http"
)

type ExpRepository struct {
	baseUrl string
	client  *http.Client
}

func NewExpRepository(baseUrl string, client *http.Client) *ExpRepository {
	return &ExpRepository{
		baseUrl: baseUrl,
		client:  client,
	}
}
