package repository

import (
	"database/sql"
	"golang-clean-server/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	query := `INSERT INTO "user" (username) VALUES ($1) RETURNING ID`
	return r.db.QueryRow(query, user.Username).Scan(&user.ID)
}
