package repository

import (
	"crypto/sha256"
	"fmt"
	"github.com/hel7/Atark-backend"
	"github.com/jmoiron/sqlx"
	"log"
)

type UserMysql struct {
	db *sqlx.DB
}

const (
	salt = "dfjaklsjlk343298hkjha"
)

func NewUserMysql(db *sqlx.DB) *UserMysql {
	return &UserMysql{db: db}
}

func (r *UserMysql) CreateUser(user farmsage.User) (int, error) {
	hashedPassword := generatePasswordHash(user.Password)

	query := fmt.Sprintf("INSERT INTO %s (username, email, password,role) VALUES (?, ?, ?, ?)", usersTable)

	result, err := r.db.Exec(query, user.Username, user.Email, hashedPassword, user.Role)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *UserMysql) GetUserByID(userID int) (farmsage.User, error) {
	var user farmsage.User
	query := fmt.Sprintf("SELECT Username, Email, Role FROM %s WHERE UserID=?", usersTable)
	err := r.db.Get(&user, query, userID)
	if err != nil {
		log.Printf("Error fetching user by ID %d: %s", userID, err)
	}
	return user, err
}

func (r *UserMysql) GetAllUsers() ([]farmsage.User, error) {
	var users []farmsage.User
	query := fmt.Sprintf("SELECT Username, Email, Role FROM %s", usersTable)
	err := r.db.Select(&users, query)
	if err != nil {
		log.Printf("Error fetching all users: %s", err)
	}
	return users, err
}
func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password + salt))
	hashed := hash.Sum(nil)
	return fmt.Sprintf("%x", hashed)
}
