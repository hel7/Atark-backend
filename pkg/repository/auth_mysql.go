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
	query := fmt.Sprintf("INSERT INTO %s (username, email, password, role) VALUES (?, ?, ?, 'User')", usersTable)
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

func (r *AuthMysql) GetUser(username, password string) (farmsage.User, error) {
	var user farmsage.User
	query := fmt.Sprintf("SELECT UserID FROM %s WHERE username=? and password=?", usersTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}
