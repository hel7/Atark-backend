package repository

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable       = "users"
	animalsTable     = "animals"
	activitiesTable  = "activities"
	feedingSchedules = "feeding_schedules"
	feedsTable       = "feeds"
	biometricsTable  = "biometrics"
	farmsTable       = "farms"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Dbname   string
}

func NewMysqlDb(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Dbname))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
