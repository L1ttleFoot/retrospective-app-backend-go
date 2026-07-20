package domain

import (
	"fmt"
	core_errors "golang-clean-server/internal/core/errors"
)

type User struct {
	ID       string
	Username string
	Password string
}

func NewUser(id string, username string, password string) User {
	return User{ID: id, Username: username, Password: password}
}

func NewUserUninit(username string, password string) User {
	return User{ID: UninitId, Username: username, Password: password}
}

func (u *User) Validate() error {
	return nil
}

type UserPatch struct {
	Username Nullable[string]
}

func (p *UserPatch) Validate() error {
	if p.Username.Set && p.Username.Value == nil {
		return fmt.Errorf("username is required: %w", core_errors.ErrInvalideArgument)
	}

	return nil
}
