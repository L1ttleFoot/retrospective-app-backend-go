package users_transport_http

import (
	"golang-clean-server/internal/core/domain"
	core_logger "golang-clean-server/internal/core/logger"
	core_request "golang-clean-server/internal/core/transport/request"
	core_response "golang-clean-server/internal/core/transport/response"
	"net/http"
)

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CreateUserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (h *UsersHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	log.Debug("invoke CreateUser handler")

	var request CreateUserRequest

	responseHandler := core_response.NewResponseHandler(log, w)

	if err := core_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate request")
		return
	}

	userDomain := domainFromDto(request)

	userDomain, err := h.usersService.CreateUser(ctx, userDomain)

	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")
		return
	}

	response := dtoFromDomain(userDomain)

	responseHandler.JsonResponse(response, http.StatusCreated)
}

func domainFromDto(dto CreateUserRequest) domain.User {
	return domain.NewUserUninit(dto.Username, dto.Password)
}

func dtoFromDomain(user domain.User) CreateUserResponse {
	return CreateUserResponse{ID: user.ID, Username: user.Username}
}
