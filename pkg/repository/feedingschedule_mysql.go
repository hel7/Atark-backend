package repository

import (
	"github.com/hel7/Atark-backend"
	"github.com/jmoiron/sqlx"
	"strings"
)

type FeedingScheduleMysql struct {
	db *sqlx.DB
}

func NewFeedingScheduleMysql(db *sqlx.DB) *FeedingScheduleMysql {
	return &FeedingScheduleMysql{db: db}
}

func (r *FeedingScheduleMysql) Create(feedingSchedule farmsage.FeedingSchedule) (int, error) {
	query := "INSERT INTO FeedingSchedule (AnimalID, FeedID, FeedingTime) VALUES (?, ?, UTC_TIMESTAMP())"
	result, err := r.db.Exec(query, feedingSchedule.AnimalID, feedingSchedule.FeedID)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastInsertID), nil
}

func (r *FeedingScheduleMysql) GetByID(animalID int) ([]farmsage.FeedingSchedule, error) {
	var feedingSchedules []farmsage.FeedingSchedule

	query := "SELECT FeedingSchedule.ScheduleID,Animal.AnimalID, Feed.FeedID, Animal.AnimalName, Animal.Number, " +
		"Feed.FeedName, FeedingSchedule.FeedingTime " +
		"FROM FeedingSchedule " +
		"INNER JOIN Animal ON FeedingSchedule.AnimalID = Animal.AnimalID " +
		"INNER JOIN Feed ON FeedingSchedule.FeedID = Feed.FeedID " +
		"WHERE FeedingSchedule.AnimalID = ?"

	err := r.db.Select(&feedingSchedules, query, animalID)
	return feedingSchedules, err
}

func (r *FeedingScheduleMysql) Delete(scheduleID int) error {
	query := "DELETE FROM FeedingSchedule WHERE ScheduleID = ?"
	_, err := r.db.Exec(query, scheduleID)
	return err
}
func (r *FeedingScheduleMysql) Update(scheduleID int, input farmsage.UpdateFeedingScheduleInput) error {
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
