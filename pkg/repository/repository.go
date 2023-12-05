package repository

import (
	farmsage "github.com/hel7/Atark-backend"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user farmsage.User) (int, error)
	GetUser(username, password string) (farmsage.User, error)
}

type Animals interface {
}

type Farms interface {
	Create(UserID int, farm farmsage.Farm) (int, error)
	GetAll(UserID int) ([]farmsage.Farm, error)
	GetByID(UserID, FarmID int) (farmsage.Farm, error)
}

type Feed interface {
}

type FeedingSchedule interface {
}

type Analytics interface {
}

type Admin interface {
	GetUserByID(userID int) (farmsage.User, error)
	CreateUser(user farmsage.User) (int, error)
	GetAllUsers() ([]farmsage.User, error)
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
		Farms:         NewFarmsMysql(db),
		Admin:         NewUserMysql(db),
	}
}
