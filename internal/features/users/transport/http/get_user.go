package users_transport_http

import (
	"golang-clean-server/internal/core/domain"
	core_logger "golang-clean-server/internal/core/logger"
	core_request "golang-clean-server/internal/core/transport/request"
	core_response "golang-clean-server/internal/core/transport/response"
	"net/http"
)

type GetUserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (h *UsersHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_response.NewResponseHandler(log, w)

	userId, err := core_request.GetPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "invalid id format")
		return
	}

	user, err := h.usersService.GetUser(ctx, userId)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user")
		return
	}

	response := dtoUserFromDomain(user)

	responseHandler.JsonResponse(response, http.StatusOK)
}

func dtoUserFromDomain(user domain.User) GetUserResponse {
	return GetUserResponse{ID: user.ID, Username: user.Username}
}
