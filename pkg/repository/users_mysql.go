package repository

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/hel7/Atark-backend"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

type UserMysql struct {
	db     *sqlx.DB
	config Config
}

func NewUserMysql(db *sqlx.DB, config Config) *UserMysql {
	return &UserMysql{db: db, config: config}
}

const (
	salt = "dfjaklsjlk343298hkjha"
)

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
func (r *UserMysql) BackupData(backupPath string) error {
	cmd := exec.Command(
		"docker",
		"exec",
		"farmsage-db",
		"mysqldump",
		"-u",
		r.config.Username,
		"-p"+r.config.Password,
		r.config.Dbname,
	)

	outputFile, err := os.Create(backupPath)
	if err != nil {
		return fmt.Errorf("error creating backup file: %s", err)
	}
	defer outputFile.Close()

	cmd.Stdout = outputFile

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error running mysqldump: %s", err)
	}

	return nil
}
func (r *UserMysql) RestoreData(backupPath string) error {
	backupPath = "backup.sql"

	containerBackupPath := "/var/lib/mysql/backup.sql"
	cmd := exec.Command(
		"docker",
		"cp",
		backupPath,
		"farmsage-db:"+containerBackupPath,
	)
	if err := cmd.Run(); err != nil {
		logrus.Errorf("Failed to copy dump file to container: %s", err)
		return err
	}

	dumpContent, err := ioutil.ReadFile(backupPath)
	if err != nil {
		logrus.Errorf("Failed to read dump file: %s", err)
		return err
	}

	mysqlCmd := exec.Command(
		"docker",
		"exec",
		"-i",
		"farmsage-db",
		"mysql",
		"-u",
		r.config.Username,
		"-p"+r.config.Password,
		r.config.Dbname,
	)

	mysqlCmd.Stdin = bytes.NewReader(dumpContent)
	mysqlCmd.Stdout = os.Stdout
	mysqlCmd.Stderr = os.Stderr

	if err := mysqlCmd.Run(); err != nil {
		logrus.Errorf("Failed to restore data in container: %s", err)
		return err
	}

	return nil
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password + salt))
	hashed := hash.Sum(nil)
	return fmt.Sprintf("%x", hashed)
}
