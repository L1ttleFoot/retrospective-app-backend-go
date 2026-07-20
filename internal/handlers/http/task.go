package handlers_http

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"golang-clean-server/internal/models"
	"net/http"
	"strconv"
)

type TaskRepository interface {
	Create(task *models.Task) error
	GetAll() ([]models.Task, error)
	GetById(id int) (models.Task, error)
	Delete(id int) (int, error)
}

type TaskHandler struct {
	repo TaskRepository
}

func NewTaskHandler(repo TaskRepository) *TaskHandler {
	return &TaskHandler{repo: repo}
}

func (h *TaskHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tasks, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newTask models.Task

	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
		return
	}

	if err := h.repo.Create(&newTask); err != nil {
		http.Error(w, "Ошибка сохранения", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	deletedID, err := h.repo.Delete(id)

	if err == sql.ErrNoRows {
		http.Error(w, "Задача с таким ID не найдена", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Ошибка сервера при удалении из БД", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("Задача с ID %d успешно удалена", deletedID),
	})
}

func (h *TaskHandler) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
