package users_transport_http

import (
	core_logger "golang-clean-server/internal/core/logger"
	core_response "golang-clean-server/internal/core/transport/response"
	"net/http"
	"time"
)

type RefreshTokensResponse struct {
	ID          string   `json:"id"`
	Username    string   `json:"username"`
	Roles       []string `json:"roles"`
	AccessToken string   `json:"token"`
}

func (h *UsersHandler) RefreshTokens(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_response.NewResponseHandler(log, w)

	cookie, err := r.Cookie("refreshToken")
	if err != nil {
		responseHandler.ErrorResponse(err, "Refresh token cookie missing")
		return
	}

	refreshToken := cookie.Value

	userDomain, err := h.usersService.RefreshTokens(ctx, refreshToken)
	if err != nil {
		responseHandler.ErrorResponse(err, "refresh tokens error")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    userDomain.RefreshToken,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   int((15 * 24 * time.Hour).Seconds()),
	})

	response := RefreshTokensResponse{
		ID:          userDomain.ID,
		Username:    userDomain.Username,
		Roles:       userDomain.Roles,
		AccessToken: userDomain.AccessToken,
	}

	responseHandler.JsonResponse(response, http.StatusOK)
}
