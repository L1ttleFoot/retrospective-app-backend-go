package exp_gateway

import (
	"context"
	"encoding/json"
	"fmt"
	core_errors "golang-clean-server/internal/core/errors"
	"io"
	"net/http"
)

type Result struct {
	Message string `json:"message"`
}

func (r *ExpRepository) ExpHttp(ctx context.Context) error {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, r.baseUrl, http.NoBody)
	if err != nil {
		return fmt.Errorf("exp http new request: %w", err)
	}

	res, err := r.client.Do(req)
	if err != nil {
		return fmt.Errorf("exp http client do: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		errBody, _ := io.ReadAll(io.LimitReader(res.Body, 4096))
		return fmt.Errorf("exp http unexpected status %d (%s): %w", res.StatusCode, string(errBody), core_errors.ErrInternal)
	}

	var result Result
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return fmt.Errorf("unmarshal response: %w", err)
	}

	fmt.Println(result)

	return nil
}
