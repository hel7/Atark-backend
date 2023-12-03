package repository

import (
	"fmt"
	farmsage "github.com/hel7/Atark-backend"
	"github.com/jmoiron/sqlx"
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

	createFarmQuery := fmt.Sprintf("INSERT INTO %s (UserID, Name) VALUES (?,?)", farmsTable)
	res, err := tx.Exec(createFarmQuery, UserID, farm.Name)

	if err != nil {
		tx.Rollback()
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
	query := fmt.Sprintf("SELECT FarmID,Name FROM %s INNER JOIN User ON Farm.UserID = User.UserID WHERE Farm.UserID = ?", farmsTable)
	err := r.db.Select(&farms, query, UserID)
	return farms, err
}

func (r *FarmsMysql) GetByID(UserID, FarmID int) (farmsage.Farm, error) {
	var farm farmsage.Farm
	query := fmt.Sprintf("SELECT FarmID,Name FROM %s INNER JOIN User ON Farm.UserID = User.UserID WHERE Farm.UserID = ? AND Farm.FarmID=?", farmsTable)
	err := r.db.Get(&farm, query, UserID, FarmID)
	return farm, err
}
