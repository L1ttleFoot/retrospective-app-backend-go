package templates_repository

import (
	"context"
	"fmt"
	core_errors "golang-clean-server/internal/core/errors"
	"time"
)

func (r *TemplatesRepository) DeleteTemplate(ctx context.Context, templateId string) error {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := `
	DELETE FROM templates
	WHERE id=$1
	`

	cmdTag, err := r.pool.Exec(ctx, query, templateId)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("user with id %s not found: %w", templateId, core_errors.ErrNotFound)
	}

	return nil
}
