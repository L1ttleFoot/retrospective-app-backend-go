package messages_transport

import (
	"context"
	"golang-clean-server/internal/core/domain"
	core_service_tokens "golang-clean-server/internal/core/service/tokens"
	core_middleware "golang-clean-server/internal/core/transport/middleware"
	core_server "golang-clean-server/internal/core/transport/server"
	"net/http"
)

type MessagesHandler struct {
	messagesService MessagesService
	tokenManager    core_service_tokens.TokenManager
}

type MessagesService interface {
	CreateMessage(ctx context.Context, message domain.Message) (domain.Message, error)
	GetMessagesBySectionId(ctx context.Context, sectionId string) ([]domain.Message, error)
	PatchMessage(ctx context.Context, messageId string, messagePatch domain.MessagePatch) (domain.Message, error)
	DeleteMessage(ctx context.Context, messageId string) error
	PinMessage(ctx context.Context, messageId string) error
	UnpinMessage(ctx context.Context, messageId string) error
	SetReaction(ctx context.Context, reaction, messageId, authorId string) error
}

func NewMessagesTransport(
	messagesService MessagesService,
	tokenManager core_service_tokens.TokenManager,
) *MessagesHandler {
	return &MessagesHandler{
		messagesService: messagesService,
		tokenManager:    tokenManager,
	}
}

func (h *MessagesHandler) Router() []core_server.Route {
	return []core_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/sections/{sectionId}/messages",
			Handler: h.GetMessagesBySectionId,
		},
		{
			Method:     http.MethodPost,
			Path:       "/messages",
			Handler:    h.CreateMessage,
			Middleware: []core_middleware.Middleware{core_middleware.Auth(h.tokenManager)},
		},
		{
			Method:  http.MethodDelete,
			Path:    "/messages/{id}",
			Handler: h.DeleteMessage,
			//Middleware: []core_middleware.Middleware{core_middleware.Auth(h.tokenManager)},
		},
		{
			Method:  http.MethodPatch,
			Path:    "/messages/{id}",
			Handler: h.PatchMessage,
			//Middleware: []core_middleware.Middleware{core_middleware.Auth(h.tokenManager)},
		},
		{
			Method:  http.MethodPost,
			Path:    "/messages/{id}/pin",
			Handler: h.PinMessage,
			//Middleware: []core_middleware.Middleware{core_middleware.Auth(h.tokenManager)},
		},
		{
			Method:  http.MethodPost,
			Path:    "/messages/{id}/unpin",
			Handler: h.UnpinMessage,
			//Middleware: []core_middleware.Middleware{core_middleware.Auth(h.tokenManager)},
		},
		{
			Method:     http.MethodPost,
			Path:       "/messages/{id}/reaction",
			Handler:    h.SetReaction,
			Middleware: []core_middleware.Middleware{core_middleware.Auth(h.tokenManager)},
		},
	}
}
