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

func userDomainsFromModels(users []UserModel) []domain.User {
	userDomains := make([]domain.User, len(users))

	for i, u := range users {
		userDomains[i] = domain.NewUser(
			u.ID, u.Username,
			"",
			"",
			"",
			nil,
		)
	}

	return userDomains
}
