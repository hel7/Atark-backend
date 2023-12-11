package repository

import (
	"github.com/hel7/Atark-backend"
	"github.com/jmoiron/sqlx"
)

type FeedingScheduleMysql struct {
	db *sqlx.DB
}

func NewFeedingScheduleMysql(db *sqlx.DB) *FeedingScheduleMysql {
	return &FeedingScheduleMysql{db: db}
}

func (r *FeedingScheduleMysql) Create(feedingSchedule farmsage.FeedingSchedule) (int, error) {
	query := "INSERT INTO FeedingSchedule (AnimalID, FeedID, FeedingTime)" +
		" VALUES (" +
		"   (SELECT AnimalID FROM Animal WHERE AnimalID = ?), " +
		"   (SELECT FeedID FROM Feed WHERE FeedID = ?), " +
		"   UTC_TIMESTAMP())"
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
	var feedingSchedule []farmsage.FeedingSchedule

	query := "SELECT Animal.AnimalName, Animal.Number , " +
		"Feed.FeedName, FeedingSchedule.FeedingTime " +
		"FROM FeedingSchedule " +
		"INNER JOIN Animal ON FeedingSchedule.AnimalID = Animal.AnimalID " +
		"INNER JOIN Feed ON FeedingSchedule.FeedID = Feed.FeedID " +
		"WHERE FeedingSchedule.AnimalID = ?"

	err := r.db.Select(&feedingSchedule, query, animalID)
	return feedingSchedule, err
}
