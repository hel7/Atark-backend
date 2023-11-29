package repository

import (
	farmsage "github.com/hel7/Atark-backend"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user farmsage.User) (int, error)
}

type Animals interface {
}

type Farms interface {
}

type Feed interface {
}

type FeedingSchedule interface {
}

type Analytics interface {
}

type Admin interface {
}

type Repository struct {
	Authorization
	Animals
	Farms
	Feed
	FeedingSchedule
	Analytics
	Admin
}

func NewRepository(db *sqlx.DB) *Repository {

	return &Repository{
		Authorization: NewAuthMysql(db),
	}
}
