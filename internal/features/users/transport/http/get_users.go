package users_transport_http

import (
	"fmt"
	"golang-clean-server/internal/core/domain"
	core_logger "golang-clean-server/internal/core/logger"
	core_response "golang-clean-server/internal/core/transport/response"
	core_http_utils "golang-clean-server/internal/core/transport/utils"
	"net/http"
)

type GetUsersResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (h *UsersHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_response.NewResponseHandler(log, w)

	limit, offset, err := gerQueryParams(r)

	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get query params")
		return
	}

	domainUsers, err := h.usersService.GetUsers(ctx, limit, offset)

	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get users")
		return
	}

	response := dtoUsersFromDomain(domainUsers)

	responseHandler.JsonResponse(response, http.StatusOK)
}

func gerQueryParams(r *http.Request) (*int, *int, error) {

	limit, err := core_http_utils.GetQueryParams(r, "limit")

	if err != nil {
		return nil, nil, fmt.Errorf("get 'limit' query param : %w", err)
	}

	offset, err := core_http_utils.GetQueryParams(r, "offset")

	if err != nil {
		return nil, nil, fmt.Errorf("get 'offset' query param : %w", err)

	}

	return limit, offset, nil
}

func dtoUsersFromDomain(users []domain.User) []GetUsersResponse {

	dto := make([]GetUsersResponse, len(users))

	for i, v := range users {
		dto[i] = GetUsersResponse{ID: v.ID, Username: v.Username}
	}

	return dto
}
