package repository

import (
	"github.com/hel7/Atark-backend"
	"github.com/jmoiron/sqlx"
)

type FeedsMysql struct {
	db *sqlx.DB
}

func NewFeedsMysql(db *sqlx.DB) *FeedsMysql {
	return &FeedsMysql{db: db}
}

func (r *FeedsMysql) Create(feed farmsage.Feed) (int, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	createFeedQuery := "INSERT INTO Feed (FeedName, Quantity) VALUES (?, ?)"
	res, err := tx.Exec(createFeedQuery, feed.FeedName, feed.Quantity)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	feedID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return 0, err
	}

	return int(feedID), nil
}

func (r *FeedsMysql) GetAll() ([]farmsage.Feed, error) {
	var feeds []farmsage.Feed
	query := "SELECT * FROM Feed"
	err := r.db.Select(&feeds, query)
	if err != nil {
		return nil, err
	}
	return feeds, nil
}

func (r *FeedsMysql) Delete(feedID int) error {
	query := "DELETE FROM Feed WHERE FeedID = ?"
	_, err := r.db.Exec(query, feedID)
	return err
}
