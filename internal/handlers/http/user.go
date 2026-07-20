package handlers_http

import (
	"encoding/json"
	"golang-clean-server/internal/models"
	"net/http"
)

type UserRepository interface {
	Create(user *models.User) error
}

type UserHandler struct {
	repo UserRepository
}

func NewUserRepository(repo UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newUser models.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
		return
	}

	if err := h.repo.Create(&newUser); err != nil {
		http.Error(w, "Ошибка сохранения", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}
