package repository

import (
	"fmt"
	farmsage "github.com/hel7/Atark-backend"
	"github.com/jmoiron/sqlx"
)

type AuthMysql struct {
	db *sqlx.DB
}

func NewAuthMysql(db *sqlx.DB) *AuthMysql {
	return &AuthMysql{db: db}
}

func (r *AuthMysql) CreateUser(user farmsage.User) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (username, email, password) VALUES (?, ?, ?)", usersTable)

	result, err := r.db.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
