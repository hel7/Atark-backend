package repository

import (
	"fmt"
	"github.com/hel7/Atark-backend"
	"github.com/jmoiron/sqlx"
)

type AnimalsMysql struct {
	db *sqlx.DB
}

func NewAnimalsMysql(db *sqlx.DB) *AnimalsMysql {
	return &AnimalsMysql{db: db}
}

func (r *AnimalsMysql) Create(UserID int, animal farmsage.Animal) (int, error) {
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

	linkAnimalQuery := "INSERT INTO Farm (UserID, AnimalID, FarmName) VALUES (?, ?, ?)"
	_, err = tx.Exec(linkAnimalQuery, UserID, animalID, "FarmName")
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
	query := fmt.Sprintf("SELECT Animal.AnimalID, Animal.AnimalName, Animal.Number, " +
		"Animal.DateOfBirth, Animal.Sex, Animal.Age, Animal.MedicalInfo " +
		"FROM Animal " +
		"WHERE Animal.AnimalID IN ( " +
		"SELECT Farm.AnimalID " +
		"FROM Farm WHERE Farm.UserID = ?)")
	err := r.db.Select(&animals, query, UserID)
	return animals, err
}

func (r *AnimalsMysql) GetByID(UserID, AnimalID int) (farmsage.Animal, error) {
	var animal farmsage.Animal
	query := fmt.Sprintf("SELECT Animal.AnimalID, Animal.AnimalName, Animal.Number, " +
		"Animal.DateOfBirth, Animal.Sex, Animal.Age, Animal.MedicalInfo " +
		"FROM Animal " +
		"INNER JOIN Farm ON Farm.AnimalID = Animal.AnimalID " +
		"WHERE Farm.UserID = ? AND Farm.AnimalID = ?")

	err := r.db.Get(&animal, query, UserID, AnimalID)
	return animal, err
}
func (r *AnimalsMysql) Delete(UserID, AnimalID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE AnimalID = ? AND EXISTS (SELECT 1 FROM Farm WHERE AnimalID = ? AND UserID = ?)", animalsTable)
	_, err := r.db.Exec(query, AnimalID, AnimalID, UserID)
	return err
}
