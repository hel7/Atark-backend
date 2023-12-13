package repository

import (
	"fmt"
	"github.com/hel7/Atark-backend"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
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

func (r *FeedsMysql) Update(feedID int, input farmsage.UpdateFeedInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)

	if input.FeedName != nil {
		setValues = append(setValues, "FeedName=?")
		args = append(args, *input.FeedName)
	}
	if input.Quantity != nil {
		setValues = append(setValues, "Quantity=?")
		args = append(args, *input.Quantity)
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE FeedID=?", feedsTable, setQuery)

	args = append(args, feedID)
	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %v", args)

	_, err := r.db.Exec(query, args...)
	return err
}
