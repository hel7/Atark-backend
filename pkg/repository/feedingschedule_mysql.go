package repository

import (
	"github.com/hel7/Atark-backend"
	"github.com/jmoiron/sqlx"
	"strings"
)

type FeedScheduleMysql struct {
	db *sqlx.DB
}

func NewFeedingScheduleMysql(db *sqlx.DB) *FeedScheduleMysql {
	return &FeedScheduleMysql{db: db}
}

func (r *FeedScheduleMysql) Create(feedingSchedule farmsage.FeedingSchedule) (int, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	query := "INSERT INTO FeedingSchedule (AnimalID, FeedID, FeedingTime, AllocatedQuantity) VALUES (?, ?, ?,?)"
	result, err := tx.Exec(query, feedingSchedule.AnimalID, feedingSchedule.FeedID, feedingSchedule.FeedingTime, feedingSchedule.AllocatedQuantity)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	getFeedQuery := "SELECT Quantity FROM Feed WHERE FeedID = ?"
	var quantity int
	err = tx.Get(&quantity, getFeedQuery, feedingSchedule.FeedID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	quantity -= feedingSchedule.AllocatedQuantity

	updateFeedQuery := "UPDATE Feed SET Quantity = ? WHERE FeedID = ?"
	_, err = tx.Exec(updateFeedQuery, quantity, feedingSchedule.FeedID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return 0, err
	}

	return int(lastInsertID), nil
}

func (r *FeedScheduleMysql) GetByID(animalID int) ([]farmsage.FeedingSchedule, error) {
	var feedingSchedules []farmsage.FeedingSchedule

	query := "SELECT FeedingSchedule.ScheduleID, Animal.AnimalID, Feed.FeedID, Animal.AnimalName, Animal.Number, " +
		"Feed.FeedName, FeedingSchedule.FeedingTime, FeedingSchedule.AllocatedQuantity " +
		"FROM FeedingSchedule " +
		"INNER JOIN Animal ON FeedingSchedule.AnimalID = Animal.AnimalID " +
		"INNER JOIN Feed ON FeedingSchedule.FeedID = Feed.FeedID " +
		"WHERE FeedingSchedule.AnimalID = ?"

	err := r.db.Select(&feedingSchedules, query, animalID)
	return feedingSchedules, err
}

func (r *FeedScheduleMysql) Delete(scheduleID int) error {
	query := "DELETE FROM FeedingSchedule WHERE ScheduleID = ?"
	_, err := r.db.Exec(query, scheduleID)
	return err
}

func (r *FeedScheduleMysql) Update(scheduleID int, input farmsage.UpdateFeedingScheduleInput) error {
	query := "UPDATE FeedingSchedule SET"

	args := make([]interface{}, 0)

	if input.AnimalID != nil {
		query += " AnimalID = ?,"
		args = append(args, *input.AnimalID)
	}

	if input.FeedID != nil {
		query += " FeedID = ?,"
		args = append(args, *input.FeedID)
	}

	if input.FeedingTime != nil {
		query += " FeedingTime = ?,"
		args = append(args, *input.FeedingTime)
	} else {
		query += " FeedingTime = NOW(),"
	}

	query = strings.TrimSuffix(query, ",")

	query += " WHERE ScheduleID = ?"
	args = append(args, scheduleID)

	_, err := r.db.Exec(query, args...)
	return err
}
