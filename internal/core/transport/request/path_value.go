package core_request

import (
	"fmt"
	core_errors "golang-clean-server/internal/core/errors"
	"net/http"

	"github.com/google/uuid"
)

func GetPathValue(r *http.Request, key string) (string, error) {

	pathValue := r.PathValue(key)

	if pathValue == "" {
		return "", fmt.Errorf(
			"no key=%s in path values : %w",
			key,
			core_errors.ErrInvalideArgument,
		)
	}

	if _, err := uuid.Parse(pathValue); err != nil {
		return "", fmt.Errorf(
			"path value='%s' by key='%s' is not a valit uuid : %v : %w",
			pathValue,
			key,
			err,
			core_errors.ErrInvalideArgument,
		)
	}

	return pathValue, nil
}
