package repository

import (
	"fmt"
	"github.com/hel7/Atark-backend"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type AnimalsMysql struct {
	db *sqlx.DB
}

func NewAnimalsMysql(db *sqlx.DB) *AnimalsMysql {
	return &AnimalsMysql{db: db}
}

func (r *AnimalsMysql) Create(FarmID int, animal farmsage.Animal) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	createAnimalQuery := "INSERT INTO Animal (AnimalName, Number, DateOfBirth, Sex, Age, MedicalInfo) VALUES (?, ?, ?, ?, ?, ?)"
	res, err := tx.Exec(createAnimalQuery, animal.AnimalName, animal.Number, animal.DateOfBirth, animal.Sex, animal.Age, animal.MedicalInfo)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	animalID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	linkAnimalQuery := "INSERT INTO FarmAnimal (FarmID, AnimalID) VALUES (?, ?)"
	_, err = tx.Exec(linkAnimalQuery, FarmID, animalID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return int(animalID), nil
}

func (r *AnimalsMysql) GetAll(UserID int) ([]farmsage.Animal, error) {
	var animals []farmsage.Animal
	query := "SELECT a.AnimalID, a.AnimalName, a.Number, a.DateOfBirth, a.Sex, a.Age, a.MedicalInfo, f.FarmName " +
		"FROM Animal AS a " +
		"JOIN FarmAnimal AS fa ON fa.AnimalID = a.AnimalID " +
		"JOIN Farm AS f ON f.FarmID = fa.FarmID " +
		"WHERE f.UserID = ?"
	err := r.db.Select(&animals, query, UserID)
	return animals, err
}

func (r *AnimalsMysql) GetByID(UserID, AnimalID int) (farmsage.Animal, error) {
	var animal farmsage.Animal
	query := "SELECT a.AnimalID, a.AnimalName, a.Number, a.DateOfBirth, a.Sex, a.Age, a.MedicalInfo, f.FarmName " +
		"FROM Animal AS a " +
		"JOIN FarmAnimal AS fa ON fa.AnimalID = a.AnimalID " +
		"JOIN Farm AS f ON f.FarmID = fa.FarmID " +
		"WHERE f.UserID = ? AND a.AnimalID = ?"
	err := r.db.Get(&animal, query, UserID, AnimalID)
	return animal, err
}

func (r *AnimalsMysql) Delete(UserID, AnimalID int) error {
	var farmID int
	err := r.db.Get(&farmID, "SELECT FarmID FROM FarmAnimal WHERE AnimalID = ?", AnimalID)
	if err != nil {
		return err
	}

	_, err = r.db.Exec("DELETE FROM FarmAnimal WHERE AnimalID = ? AND FarmID = ?", AnimalID, farmID)
	if err != nil {
		return err
	}

	_, err = r.db.Exec("DELETE FROM Animal WHERE AnimalID = ?", AnimalID)
	return err
}

func (r *AnimalsMysql) Update(AnimalID int, input farmsage.UpdateAnimalInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	setValues := make([]string, 0)
	args := make([]interface{}, 0)

	if input.AnimalName != nil {
		setValues = append(setValues, "AnimalName=?")
		args = append(args, *input.AnimalName)
	}
	if input.Number != nil {
		setValues = append(setValues, "Number=?")
		args = append(args, *input.Number)
	}
	if input.DateOfBirth != nil {
		setValues = append(setValues, "DateOfBirth=?")
		args = append(args, *input.DateOfBirth)
	}
	if input.Sex != nil {
		setValues = append(setValues, "Sex=?")
		args = append(args, *input.Sex)
	}
	if input.Age != nil {
		setValues = append(setValues, "Age=?")
		args = append(args, *input.Age)
	}
	if input.MedicalInfo != nil {
		setValues = append(setValues, "MedicalInfo=?")
		args = append(args, *input.MedicalInfo)
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE AnimalID=?", animalsTable, setQuery)

	args = append(args, AnimalID)
	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %v", args)

	_, err := r.db.Exec(query, args...)
	return err
}
