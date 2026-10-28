package users_transport_http

import (
	"golang-clean-server/internal/core/domain"
	core_logger "golang-clean-server/internal/core/logger"
	core_request "golang-clean-server/internal/core/transport/request"
	core_response "golang-clean-server/internal/core/transport/response"
	"net/http"
)

type LoginUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginUserResponse struct {
	ID          string   `json:"id"`
	Username    string   `json:"username"`
	AccessToken string   `json:"token"`
	Roles       []string `json:"roles"`
}

func (h *UsersHandler) LoginUser(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	var request LoginUserRequest

	responseHandler := core_response.NewResponseHandler(log, w)

	if err := core_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate request")
		return
	}

	userDomain := domain.User{ID: "", Username: request.Username, Password: request.Password}

	userDomain, err := h.usersService.LoginUser(ctx, userDomain)

	if err != nil {
		responseHandler.ErrorResponse(err, "failed to login")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    userDomain.RefreshToken,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	response := LoginUserResponse{
		ID:          userDomain.ID,
		Username:    userDomain.Username,
		Roles:       userDomain.Roles,
		AccessToken: userDomain.AccessToken,
	}

	responseHandler.JsonResponse(response, http.StatusOK)
}
