package repositories

import (
	"api/src/models"
	"database/sql"
)

// UsersRepository - This is the struct that will contain the connection to the database
type UsersRepository struct {
	db *sql.DB
}

// NewUsersRepository - This is the function that will receive the connection to the database and return the struct
func NewUsersRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{db}
}

// Create - This is a method of the struct users
func (repository UsersRepository) Create(user models.User) (uint64, error) {
	// prepare the statement to save user
	statement, err := repository.db.Prepare(
		"insert into users (name, nick, email, password) values (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	// execute the statement and save user
	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	// return the last id created
	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastID), nil
}

func (repository UsersRepository) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = "%" + nameOrNick + "%"

	lines, err := repository.db.Query(
		"select id, name, nick, email, createdAt from users where name LIKE ? or nick LIKE ?", nameOrNick, nameOrNick,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User
	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil

}
