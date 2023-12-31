package repository

import (
	farmsage "github.com/hel7/Atark-backend"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user farmsage.User) (int, error)
	GetUser(username, password string) (farmsage.User, error)
	CreateAdmin(user farmsage.User) (int, error)
}

type Animals interface {
	Create(FarmID int, animal farmsage.Animal) (int, error)
	GetAll(FarmID int) ([]farmsage.Animal, error)
	GetByID(FarmID, AnimalID int) (farmsage.Animal, error)
	Delete(FarmID, AnimalID int) error
	Update(AnimalID int, input farmsage.UpdateAnimalInput) error
	AddActivity(AnimalID int, activity farmsage.Activity) (int, error)
	GetActivityByAnimalID(AnimalID int) ([]farmsage.Activity, error)
	AddBiometrics(AnimalID int, biometrics farmsage.Biometrics) (int, error)
	GetBiometricsByAnimalID(AnimalID int) ([]farmsage.Biometrics, error)
	DeleteBiometrics(BiometricID int) error
	DeleteActivity(ActivityID int) error
}

type Farms interface {
	Create(UserID int, farm farmsage.Farm) (int, error)
	GetAll(UserID int) ([]farmsage.Farm, error)
	GetByID(UserID, FarmID int) (farmsage.Farm, error)
	Delete(UserID, FarmID int) error
	Update(UserID, id int, input farmsage.UpdateFarmInput) error
}

type Feed interface {
	Create(feed farmsage.Feed) (int, error)
	GetAll() ([]farmsage.Feed, error)
	Delete(feedID int) error
	Update(id int, input farmsage.UpdateFeedInput) error
}

type FeedingSchedule interface {
	Create(feedingSchedule farmsage.FeedingSchedule) (int, error)
	GetByID(animalID int) ([]farmsage.FeedingSchedule, error)
	Delete(scheduleID int) error
	Update(scheduleID int, input farmsage.UpdateFeedingScheduleInput) error
}

type Admin interface {
	GetUserByID(userID int) (farmsage.User, error)
	CreateUser(user farmsage.User) (int, error)
	GetAllUsers() ([]farmsage.User, error)
	Delete(UserID int) error
	UpdateUser(UserID int, input farmsage.UpdateUserInput, user farmsage.User) error
	BackupData(backupPath string) error
	RestoreData(backupPath string) error
	ExportData(exportPath string) error
	ImportData(importPath string) error
}

type Repository struct {
	Authorization
	Animals
	Farms
	Feed
	FeedingSchedule
	Admin
}

func NewRepository(db *sqlx.DB, config Config) *Repository {
	return &Repository{
		Authorization:   NewAuthMysql(db),
		Animals:         NewAnimalsMysql(db),
		Farms:           NewFarmsMysql(db),
		Feed:            NewFeedsMysql(db),
		FeedingSchedule: NewFeedingScheduleMysql(db),
		Admin:           NewAdminMysql(db, config),
	}
}
