package repository

import (
	"crypto/sha256"
	"fmt"
	"github.com/hel7/Atark-backend"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"log"
	"strings"
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

func (r *UserMysql) Delete(UserID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE UserID = ? ", usersTable)
	_, err := r.db.Exec(query, UserID)
	return err
}
func (r *UserMysql) UpdateUser(UserID int, input farmsage.UpdateUserInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	setValues := make([]string, 0)
	args := make([]interface{}, 0)

	if input.Username != nil {
		setValues = append(setValues, "Username=?")
		args = append(args, *input.Username)
	}
	if input.Email != nil {
		setValues = append(setValues, "Email=?")
		args = append(args, *input.Email)
	}
	if input.Password != nil {
		setValues = append(setValues, "Password=?")
		args = append(args, *input.Password)
	}
	if input.Role != nil {
		setValues = append(setValues, "Role=?")
		args = append(args, *input.Role)
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE UserID=?", usersTable, setQuery)

	args = append(args, UserID)
	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %v", args)

	_, err := r.db.Exec(query, args...)
	return err
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password + salt))
	hashed := hash.Sum(nil)
	return fmt.Sprintf("%x", hashed)
}
