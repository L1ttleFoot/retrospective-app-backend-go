package core_sse

type Event struct {
	BoardID string
	Type    string
	Data    []byte
}
