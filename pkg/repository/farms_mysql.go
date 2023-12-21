package repository

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	farmsage "github.com/hel7/Atark-backend"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type FarmsMysql struct {
	db *sqlx.DB
}

func NewFarmsMysql(db *sqlx.DB) *FarmsMysql {
	return &FarmsMysql{db: db}
}

func (r *FarmsMysql) Create(UserID int, farm farmsage.Farm) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	createFarmQuery := fmt.Sprintf("INSERT INTO %s (UserID, FarmName) VALUES (?, ?)", farmsTable)
	res, err := tx.Exec(createFarmQuery, UserID, farm.FarmName)
	if err != nil {
		tx.Rollback()

		// Проверяем, является ли ошибка ошибкой дубликата уникального ключа
		mysqlErr, ok := err.(*mysql.MySQLError)
		if ok && mysqlErr.Number == 1062 {
			return 0, fmt.Errorf("farm with this name already exists")
		}

		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return int(id), nil
}

func (r *FarmsMysql) GetAll(UserID int) ([]farmsage.Farm, error) {
	var farms []farmsage.Farm
	query := fmt.Sprintf("SELECT FarmID,FarmName FROM %s INNER JOIN User ON Farm.UserID = User.UserID WHERE Farm.UserID = ?", farmsTable)
	err := r.db.Select(&farms, query, UserID)
	return farms, err
}

func (r *FarmsMysql) GetByID(UserID, FarmID int) (farmsage.Farm, error) {
	var farm farmsage.Farm
	query := fmt.Sprintf("SELECT FarmID,FarmName FROM %s INNER JOIN User ON Farm.UserID = User.UserID WHERE Farm.UserID = ? AND Farm.FarmID=?", farmsTable)
	err := r.db.Get(&farm, query, UserID, FarmID)
	return farm, err
}
func (r *FarmsMysql) Delete(UserID, FarmID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE FarmID = ? AND UserID = ? ", farmsTable)
	_, err := r.db.Exec(query, FarmID, UserID)
	return err
}

func (r *FarmsMysql) Update(UserID, id int, input farmsage.UpdateFarmInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)

	if input.FarmName != nil {
		setValues = append(setValues, "FarmName=?")
		args = append(args, *input.FarmName)
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE FarmID=? AND UserID=?", farmsTable, setQuery)

	args = append(args, id, UserID)
	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %v", args)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		mysqlErr, ok := err.(*mysql.MySQLError)
		if ok && mysqlErr.Number == 1062 {
			return fmt.Errorf("farm with this name already exists")
		}
		return err
	}
	return nil
}
