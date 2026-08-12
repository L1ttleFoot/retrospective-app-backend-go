package events

import (
	"fmt"
	core_logger "golang-clean-server/internal/core/logger"
	core_server "golang-clean-server/internal/core/transport/server"
	"net/http"
	"time"
)

type SSE struct {
}

func (b *Broker) StreamEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Transfer-Encoding", "chunked")

	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	log.Debug("новый клиент")

	flusher, ok := w.(http.Flusher)
	if ok {
		// Отправляем пустой хук или:ok сразу, чтобы "пробить" соединение до браузера
		fmt.Fprintf(w, ":ok\n\n")
		flusher.Flush()
	}

	messageChan := make(chan []byte, 1)

	b.newClients <- messageChan

	defer func() {
		b.closingClients <- messageChan
	}()

	pingTicker := time.NewTicker(15 * time.Second)
	defer pingTicker.Stop()

	flusher, hasFlusher := w.(http.Flusher)

	for {
		select {
		case <-ctx.Done(): // Если клиент закрыл вкладку/соединение
			log.Debug("Клиент отключился")
			return
		case <-pingTicker.C:
			// Отправляем SSE-комментарий в качестве пинга
			// Браузер его проигнорирует, но прокси-сервер поймет, что соединение живо
			_, err := fmt.Fprintf(w, ":ok\n\n")
			if err != nil {
				return
			}
			if hasFlusher {
				flusher.Flush()
			}
		case msg, open := <-messageChan: // Ждем новое сообщение
			if !open {
				return
			}
			// Форматирование сообщения по спецификации SSE (data: <содержимое>\n\n)
			_, err := fmt.Fprintf(w, "data: %s\n\n", msg)
			if err != nil {
				return
			}

			// Сброс буфера в сеть, чтобы клиент сразу получил данные
			if flusher, ok := w.(http.Flusher); ok {
				flusher.Flush()
			}
		}
	}
}

func (b *Broker) Routes() []core_server.Route {
	return []core_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/event/",
			Handler: b.StreamEvent,
		},
	}
}
