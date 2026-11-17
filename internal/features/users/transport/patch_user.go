package users_transport

import (
	"golang-clean-server/internal/core/domain"
	core_logger "golang-clean-server/internal/core/logger"
	core_request "golang-clean-server/internal/core/transport/request"
	core_response "golang-clean-server/internal/core/transport/response"
	core_types "golang-clean-server/internal/core/transport/types"
	"net/http"
)

type ParchUserRequest struct {
	Username core_types.Nullable[string] `json:"username"`
}

type PatchUserRespose struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (h *UsersHandler) PatchUser(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_response.NewResponseHandler(log, w)

	userId, err := core_request.GetPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "invalid id format")
		return
	}

	var request ParchUserRequest

	if err := core_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate request")
		return
	}

	userPatch := userPatchFromRequest(request)

	userDomain, err := h.usersService.PatchUser(ctx, userId, userPatch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed tp patch user")
		return
	}

	response := PatchUserRespose{ID: userDomain.ID, Username: userDomain.Username}

	responseHandler.JsonResponse(response, http.StatusOK)
}

func userPatchFromRequest(request ParchUserRequest) domain.UserPatch {
	return domain.NewUserPatch(request.Username.ToDomain())
}
