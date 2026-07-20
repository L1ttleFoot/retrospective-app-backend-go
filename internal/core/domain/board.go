package domain

type Board struct {
	ID         string
	Title      string
	BoardOwner string
}

func NewBoard(id string, title string, boardOwner string) Board {
	return Board{ID: id, Title: title, BoardOwner: boardOwner}
}
