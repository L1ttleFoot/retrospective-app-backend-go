package sections_transport

import "net/http"

type SectionsItem struct {
	Title   string `json:"title"`
	Color   string `json:"color"`
	BoardId string `json:"boardId"`
}

type CreateSectionsRequest struct {
	Sections []SectionsItem `json:"sections"`
}

func (h *SectionsHandler) CreateSections(w http.ResponseWriter, r *http.Request) {

}
