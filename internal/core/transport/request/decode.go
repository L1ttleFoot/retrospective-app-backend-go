package core_request

import (
	"encoding/json"
	"fmt"
	core_errors "golang-clean-server/internal/core/errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var requestValidator = validator.New()

func DecodeAndValidate(r *http.Request, dto any) error {
	if err := json.NewDecoder(r.Body).Decode(dto); err != nil {
		return fmt.Errorf(
			"decode json: %v: %w",
			err,
			core_errors.ErrInvalideArgument,
		)

	}

	if err := requestValidator.Struct(dto); err != nil {
		return fmt.Errorf(
			"request validation: %v: %w",
			err,
			core_errors.ErrInvalideArgument,
		)
	}

	return nil
}
