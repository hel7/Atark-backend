package repository

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/hel7/Atark-backend"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

type AdminMysql struct {
	db     *sqlx.DB
	config Config
}

func NewAdminMysql(db *sqlx.DB, config Config) *AdminMysql {
	return &AdminMysql{db: db, config: config}
}

const (
	salt = "dfjaklsjlk343298hkjha"
)

func (r *AdminMysql) CreateUser(user farmsage.User) (int, error) {
	checkQuery := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE email=? OR username=?", usersTable)
	var count int
	err := r.db.Get(&count, checkQuery, user.Email, user.Username)
	if err != nil {
		return 0, err
	}

	if count > 0 {
		return 0, fmt.Errorf("user with this email or username already exists")
	}

	hashedPassword := generatePasswordHash(user.Password)

	query := fmt.Sprintf("INSERT INTO %s (username, email, password, role) VALUES (?, ?, ?, ?)", usersTable)

	result, err := r.db.Exec(query, user.Username, user.Email, hashedPassword, user.Role)
	if err != nil {
		mysqlErr, ok := err.(*mysql.MySQLError)
		if ok && mysqlErr.Number == 1062 {
			return 0, fmt.Errorf("user with this email or username already exists")
		}
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *AdminMysql) GetUserByID(userID int) (farmsage.User, error) {
	var user farmsage.User
	query := fmt.Sprintf("SELECT Username, Email, Role FROM %s WHERE UserID=?", usersTable)
	err := r.db.Get(&user, query, userID)
	if err != nil {
		log.Printf("Error fetching user by ID %d: %s", userID, err)
	}
	return user, err
}

func (r *AdminMysql) GetAllUsers() ([]farmsage.User, error) {
	var users []farmsage.User
	query := fmt.Sprintf("SELECT Username, Email, Role FROM %s", usersTable)
	err := r.db.Select(&users, query)
	if err != nil {
		log.Printf("Error fetching all users: %s", err)
	}
	return users, err
}

func (r *AdminMysql) Delete(UserID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE UserID = ? ", usersTable)
	_, err := r.db.Exec(query, UserID)
	return err
}

func (r *AdminMysql) UpdateUser(UserID int, input farmsage.UpdateUserInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	setValues := make([]string, 0)
	args := make([]interface{}, 0)

	if input.Username != nil {
		checkUsernameQuery := "SELECT COUNT(*) FROM User WHERE Username=? AND UserID <> ?"
		var usernameCount int
		err := r.db.Get(&usernameCount, checkUsernameQuery, *input.Username, UserID)
		if err != nil {
			return err
		}

		if usernameCount > 0 {
			return fmt.Errorf("user with this username already exists")
		}

		setValues = append(setValues, "Username=?")
		args = append(args, *input.Username)
	}
	if input.Email != nil {
		// Перевірка, чи існує користувач з такою ж поштою
		checkEmailQuery := "SELECT COUNT(*) FROM User WHERE Email=? AND UserID <> ?"
		var emailCount int
		err := r.db.Get(&emailCount, checkEmailQuery, *input.Email, UserID)
		if err != nil {
			return err
		}

		if emailCount > 0 {
			return fmt.Errorf("user with this email already exists")
		}

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

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *AdminMysql) BackupData(backupPath string) error {
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
func (r *AdminMysql) RestoreData(backupPath string) error {
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
func (r *AdminMysql) ExportData(exportPath string) error {
	logger := logrus.New()

	file := xlsx.NewFile()

	tables := map[string]string{
		"User":            "SELECT * FROM User;",
		"Feed":            "SELECT * FROM Feed;",
		"Animal":          "SELECT * FROM Animal;",
		"Farm":            "SELECT * FROM Farm;",
		"Activity":        "SELECT * FROM Activity;",
		"Biometrics":      "SELECT * FROM Biometrics;",
		"FeedingSchedule": "SELECT * FROM FeedingSchedule;",
		"FarmAnimal":      "SELECT * FROM FarmAnimal;",
	}

	for tableName, query := range tables {
		rows, err := r.db.Query(query)
		if err != nil {
			logger.Errorf("Error querying table %s: %s", tableName, err)
			return fmt.Errorf("error querying table %s: %w", tableName, err)
		}
		defer rows.Close()

		columns, err := rows.Columns()
		if err != nil {
			logger.Errorf("Error fetching columns for table %s: %s", tableName, err)
			return fmt.Errorf("error fetching columns for table %s: %w", tableName, err)
		}

		sheet, err := file.AddSheet(tableName)
		if err != nil {
			logger.Errorf("Error adding sheet %s: %s", tableName, err)
			return fmt.Errorf("error adding sheet %s: %w", tableName, err)
		}

		headerRow := sheet.AddRow()
		for _, column := range columns {
			cell := headerRow.AddCell()
			cell.Value = column
		}

		for rows.Next() {
			rowData := make([]string, len(columns))
			valuePointers := make([]interface{}, len(columns))
			for i := range rowData {
				valuePointers[i] = &rowData[i]
			}

			err := rows.Scan(valuePointers...)
			if err != nil {
				logger.Errorf("Error scanning rows for table %s: %s", tableName, err)
				return fmt.Errorf("error scanning rows for table %s: %w", tableName, err)
			}

			for i := range rowData {
				if bytesVal, ok := valuePointers[i].(*[]byte); ok {
					rowData[i] = string(*bytesVal)
				}
			}

			row := sheet.AddRow()
			for _, value := range rowData {
				cell := row.AddCell()
				cell.Value = value
			}
		}
	}

	err := file.Save(exportPath)
	if err != nil {
		logger.Errorf("Error saving Excel file: %s", err)
		return fmt.Errorf("error saving Excel file: %w", err)
	}

	return nil
}
func (r *AdminMysql) ImportData(importPath string) error {
	file, err := xlsx.OpenFile(importPath)
	if err != nil {
		logrus.Errorf("Error opening Excel file: %s", err)
		return fmt.Errorf("error opening Excel file: %w", err)
	}

	for _, sheet := range file.Sheets {
		tableName := sheet.Name

		rows := sheet.Rows
		if len(rows) < 2 {
			continue
		}

		columns := make([]string, len(rows[0].Cells))
		for i, cell := range rows[0].Cells {
			columns[i] = cell.String()
		}

		query := fmt.Sprintf("INSERT INTO %s (%s) VALUES ", tableName, strings.Join(columns, ","))

		var valueStrings []string
		var valueArgs []interface{}

		for _, row := range rows[1:] {
			var values []interface{}

			for _, cell := range row.Cells {
				values = append(values, cell.Value)
			}

			placeholders := make([]string, len(columns))
			for i := range placeholders {
				placeholders[i] = "?"
			}

			valueStrings = append(valueStrings, "("+strings.Join(placeholders, ",")+")")
			valueArgs = append(valueArgs, values...)
		}

		query += strings.Join(valueStrings, ",")
		_, err := r.db.Exec(query, valueArgs...)
		if err != nil {
			logrus.Errorf("Error inserting data into table %s: %s", tableName, err)
			return fmt.Errorf("error inserting data into table %s: %w", tableName, err)
		}
	}

	return nil
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password + salt))
	hashed := hash.Sum(nil)
	return fmt.Sprintf("%x", hashed)
}
