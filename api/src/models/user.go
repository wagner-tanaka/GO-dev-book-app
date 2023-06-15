package models

import (
	"errors"
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

func (user *User) Prepare() error {
	if err := user.validate(); err != nil {
		return err
	}

	user.format()

	return nil
}
func (user *User) validate() error {
	if user.Name == "" {
		return errors.New(" The name is required and can not be empty")
	}

	if user.Nick == "" {
		return errors.New(" The nick is required and can not be empty")
	}

	if user.Email == "" {
		return errors.New(" The email is required and can not be empty")
	}

	if user.Password == "" {
		return errors.New(" The password is required and can not be empty")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
