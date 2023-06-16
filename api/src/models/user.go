package models

import (
	"api/src/security"
	"errors"
	"github.com/badoux/checkmail"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
}

func (user *User) Prepare(stage string) error {
	if err := user.validate(stage); err != nil {
		return err
	}

	if err := user.format(stage); err != nil {
		return err
	}

	return nil
}
func (user *User) validate(stage string) error {
	if user.Name == "" {
		return errors.New(" The name is required and can not be empty")
	}

	if user.Nick == "" {
		return errors.New(" The nick is required and can not be empty")
	}

	if user.Email == "" {
		return errors.New(" The email is required and can not be empty")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("email invalid")
	}

	if stage == "register" && user.Password == "" {
		return errors.New(" The password is required and can not be empty")
	}

	return nil
}

func (user *User) format(stage string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if stage == "register" {
		passwordWithHash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(passwordWithHash)
	}

	return nil
}
