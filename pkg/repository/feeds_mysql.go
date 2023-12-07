package repository

import (
	"fmt"
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

	createFeedQuery := "INSERT INTO Feed (Name, Quantity) VALUES (?, ?)"
	res, err := tx.Exec(createFeedQuery, feed.Name, feed.Quantity)
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
	query := fmt.Sprintf("SELECT * FROM %s", feedsTable)
	err := r.db.Select(&feeds, query)
	if err != nil {
		return nil, err
	}
	return feeds, nil
}
