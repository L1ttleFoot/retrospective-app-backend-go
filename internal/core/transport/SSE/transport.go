package core_sse

import (
	"fmt"
	core_logger "golang-clean-server/internal/core/logger"
	core_server "golang-clean-server/internal/core/transport/server"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func (b *Broker) StreamEvent(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unsupported", http.StatusInternalServerError)
		return
	}

	boardID := r.PathValue("boardId")

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	log.Debug("новый клиент", zap.String("boardId", boardID))

	fmt.Fprintf(w, ":ok\n\n")
	flusher.Flush()

	messageChan := make(chan Event, 16)

	b.newClients <- client{ch: messageChan, boardID: boardID}

	defer func() {
		b.closingClients <- messageChan
	}()

	pingTicker := time.NewTicker(15 * time.Second)
	defer pingTicker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Debug("Клиент отключился")
			return
		case <-pingTicker.C:
			if _, err := fmt.Fprintf(w, ":ok\n\n"); err != nil {
				return
			}
			flusher.Flush()
		case event, open := <-messageChan:
			if !open {
				return
			}
			if _, err := fmt.Fprintf(w, "event: %s\ndata: %s\n\n", event.Type, event.Data); err != nil {
				return
			}
			flusher.Flush()
		}
	}
}

func (b *Broker) Routes() []core_server.Route {
	return []core_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/events/boards/{boardId}",
			Handler: b.StreamEvent,
		},
	}
}
