package accounts

import (
	"database/sql"
	"errors"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

const USERS_TABLE = "users"

type UsersRepository struct {
	db *utils.DatabaseAdapdater
}

func (r UsersRepository) GetById(userId string) (User, error) {
	scanner := r.db.SelectOne("id, username, name, phone_number, created_at, updated_at", USERS_TABLE, squirrel.Eq{"id": userId})
	user, err := buildUserFromDb(*scanner)

	if err != nil {
		log.Fatalf(err.Error())
		return User{}, err
	}

	return user, nil
}

func (r UsersRepository) GetAllSubscribed() []User {
	rows, err := r.db.Select("id, username, name, phone_number, created_at, updated_at", USERS_TABLE, squirrel.Eq{"subscribed": 1})
	defer rows.Close()

	if err != nil {
		return make([]User, 0)
	}

	return buildUsersFromDb(rows)
}

func NewUsersRepository() UsersRepository {
	return UsersRepository{
		utils.GetDatabase(),
	}
}

func buildUserFromDb(row squirrel.RowScanner) (User, error) {
	var user User

	err := row.Scan(
		&user.Id,
		&user.Username,
		&user.Name,
		&user.PhoneNumber,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return User{}, err
	}

	if user.Id.String() == "00000000-0000-0000-0000-000000000000" {
		return User{}, errors.New("User not found")
	}

	return user, nil
}

func buildUsersFromDb(rows *sql.Rows) []User {
	users := make([]User, 0)

	for rows.Next() {
		var user User
		rows.Scan(
			&user.Id,
			&user.Username,
			&user.Name,
			&user.PhoneNumber,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		users = append(users, user)
	}

	return users
}
