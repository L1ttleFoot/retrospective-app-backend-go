package core_response

import (
	"encoding/json"
	"errors"
	"fmt"
	core_errors "golang-clean-server/internal/core/errors"
	core_logger "golang-clean-server/internal/core/logger"
	"net/http"

	"go.uber.org/zap"
)

type ResponseHandler struct {
	log *core_logger.Logger
	w   http.ResponseWriter
}

func NewResponseHandler(log *core_logger.Logger, w http.ResponseWriter) *ResponseHandler {
	return &ResponseHandler{log: log, w: w}
}

func (h *ResponseHandler) JsonResponse(responseBody any, statusCode int) {

	h.w.WriteHeader(statusCode)

	if err := json.NewEncoder(h.w).Encode(responseBody); err != nil {
		h.log.Error("write http response", zap.Error(err))
	}

}

func (h *ResponseHandler) ErrorResponse(err error, message string) {

	var statusCode int
	var clientErrorText string

	switch {
	case errors.Is(err, core_errors.ErrUnauthorized):
		statusCode = http.StatusUnauthorized
		clientErrorText = "Не авторизован"
	case errors.Is(err, core_errors.ErrForbidden):
		statusCode = http.StatusForbidden
		clientErrorText = "Доступ запрещен"
	case errors.Is(err, core_errors.ErrInvalidArgument):
		statusCode = http.StatusBadRequest
		clientErrorText = "Неверные параметры запроса"
	case errors.Is(err, core_errors.ErrNotFound):
		statusCode = http.StatusNotFound
		clientErrorText = "Ресурс не найден"
	default:
		statusCode = http.StatusInternalServerError
		clientErrorText = "Внутренняя ошибка сервера"
		message = "Произошла непредвиденная ошибка. Пожалуйста, попробуйте позже."
	}

	if statusCode == http.StatusInternalServerError {
		h.log.Error("internal server error occurred", zap.Error(err), zap.String("context_msg", message))
	} else {
		h.log.Warn("client error occurred", zap.Error(err), zap.String("context_msg", message), zap.Int("status", statusCode))
	}

	h.w.WriteHeader(statusCode)

	response := map[string]string{"message": message, "error": clientErrorText}

	if err := json.NewEncoder(h.w).Encode(response); err != nil {
		h.log.Error("write http response", zap.Error(err))
	}
}

func (h *ResponseHandler) PanicResponse(p any, msg string) {
	statusCode := http.StatusInternalServerError

	err := fmt.Errorf("unexpected panic %v", p)

	h.log.Error(msg, zap.Error(err))

	h.w.WriteHeader(statusCode)

	response := map[string]string{"message": msg, "error": err.Error()}

	if err := json.NewEncoder(h.w).Encode(response); err != nil {
		h.log.Error("write http respons", zap.Error(err))
	}
}

func (h *ResponseHandler) EmptyResponse(statusCode int) {
	h.w.WriteHeader(statusCode)
}
