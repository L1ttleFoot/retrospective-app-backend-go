package repository

import (
	"database/sql"
	"golang-clean-server/internal/models"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetAll() ([]models.Task, error) {

	query := "SELECT id, title, done, user_id FROM tasks"

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasks := make([]models.Task, 0)

	for rows.Next() {
		var t models.Task

		if err := rows.Scan(&t.ID, &t.Title, &t.Done, &t.User_id); err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (r *TaskRepository) Create(task *models.Task) error {
	query := "INSERT INTO tasks (title, done, user_id) VALUES ($1, $2, $3) RETURNING id"
	return r.db.QueryRow(query, task.Title, task.Done, task.User_id).Scan(&task.ID)
}

func (r *TaskRepository) Delete(id int) (int, error) {

	query := "DELETE FROM tasks WHERE id = $1 RETURNING id"

	var deletedID int

	err := r.db.QueryRow(query, id).Scan(&deletedID)

	return deletedID, err

}

func (r *TaskRepository) GetById(id int) (models.Task, error) {

	return models.Task{}, nil
}
