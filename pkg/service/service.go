package service

import (
	farmsage "github.com/hel7/Atark-backend"
	"github.com/hel7/Atark-backend/pkg/repository"
)

type Authorization interface {
	CreateUser(user farmsage.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Animals interface {
	Create(UserID int, animal farmsage.Animal) (int, error)
	GetAll(UserID int) ([]farmsage.Animal, error)
	GetByID(UserID, AnimalID int) (farmsage.Animal, error)
	Delete(UserID, AnimalID int) error
}

type Farms interface {
	Create(UserID int, farm farmsage.Farm) (int, error)
	GetAll(UserID int) ([]farmsage.Farm, error)
	GetByID(UserID, FarmID int) (farmsage.Farm, error)
	Delete(UserID, FarmID int) error
}

type Feed interface {
	Create(feed farmsage.Feed) (int, error)
	GetAll() ([]farmsage.Feed, error)
	Delete(feedID int) error
}

type FeedingSchedule interface {
	Create(feedingSchedule farmsage.FeedingSchedule) (int, error)
	GetByID(animalID int) ([]farmsage.FeedingSchedule, error)
	Delete(scheduleID int) error
}

type Analytics interface {
}

type Admin interface {
	GetUserByID(userID int) (farmsage.User, error)
	GetAllUsers() ([]farmsage.User, error)
	CreateUser(user farmsage.User) (int, error)
}

type Service struct {
	Authorization
	Animals
	Farms
	Feed
	FeedingSchedule
	Analytics
	Admin
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization:   NewAuthService(repos.Authorization),
		Farms:           NewFarmService(repos.Farms),
		Admin:           NewAdminService(repos.Admin),
		Animals:         NewAnimalService(repos.Animals),
		Feed:            NewFeedService(repos.Feed),
		FeedingSchedule: NewFeedingScheduleService(repos.FeedingSchedule),
	}
}
