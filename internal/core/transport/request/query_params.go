package core_request

import (
	"fmt"
	core_errors "golang-clean-server/internal/core/errors"
	"net/http"
	"strconv"
)

func GetQueryParams[T any](r *http.Request, key string) (*T, error) {
	param := r.URL.Query().Get(key)

	if param == "" {
		return nil, nil
	}

	var (
		result T
		err    error
	)

	switch any(&result).(type) {
	case *string:
		*any(&result).(*string) = param
	case *int:
		var tmp int
		tmp, err = strconv.Atoi(param)
		if err == nil {
			*any(&result).(*int) = tmp
		}
	case *bool:
		var tmp bool
		tmp, err = strconv.ParseBool(param)
		if err == nil {
			*any(&result).(*bool) = tmp
		}
	default:
		return nil, fmt.Errorf("unsupported type for query param key='%s'", key)
	}

	if err != nil {
		return nil, fmt.Errorf(
			"param='%s' by key='%s' is not a valid value : %v : %w",
			param, key, err, core_errors.ErrInvalideArgument,
		)
	}

	return &result, nil
}
