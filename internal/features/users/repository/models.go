package user_repository

import "golang-clean-server/internal/core/domain"

type UserModel struct {
	ID       string
	Username string
	Password string
}

func (m UserModel) userModelToDomain() domain.User {
	return domain.NewUser(
		m.ID,
		m.Username,
		"",
		"",
		"",
		nil,
	)
}

