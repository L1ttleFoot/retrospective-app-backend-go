package domain

import (
	"fmt"
	core_errors "golang-clean-server/internal/core/errors"
)

type User struct {
	ID           string
	Username     string
	Password     string
	AccessToken  string
	RefreshToken string
	Roles        []string
}

func NewUser(
	id string,
	username string,
	password string,
	accessToken string,
	refreshToken string,
	roles []string,
) User {
	return User{
		ID:           id,
		Username:     username,
		Password:     password,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Roles:        roles,
	}
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

func NewUserPatch(username Nullable[string]) UserPatch {
	return UserPatch{Username: username}
}

func (p *UserPatch) Validate() error {
	if p.Username.Set && p.Username.Value == nil {
		return fmt.Errorf("username is required: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}
