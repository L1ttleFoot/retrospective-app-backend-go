package core_http_utils

import (
	"fmt"
	core_errors "golang-clean-server/internal/core/errors"
	"net/http"
	"strconv"
)

func GetQueryParams(r *http.Request, key string) (*int, error) {
	param := r.URL.Query().Get(key)

	if param == "" {
		return nil, nil
	}

	val, err := strconv.Atoi(param)

	if err != nil {
		return nil, fmt.Errorf(
			"param='%s' by key='%s' is not a valit integer : %v : %w",
			param, key, err, core_errors.ErrInvalideArgument,
		)
	}

	return &val, nil
}

func GetQueryParamsString(r *http.Request, key string) (*string, error) {
	param := r.URL.Query().Get(key)

	if param == "" {
		return nil, nil
	}

	return &param, nil
}
