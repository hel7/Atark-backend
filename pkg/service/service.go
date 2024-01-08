package service

import (
	farmsage "github.com/hel7/Atark-backend"
	"github.com/hel7/Atark-backend/pkg/repository"
)

// Authorization описує методи, необхідні для авторизації користувачів та генерації токенів.
type Authorization interface {
	CreateAdmin(user farmsage.User) (int, error)
	CreateUser(user farmsage.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

// Animals визначає методи, які пов'язані з операціями над тваринами у фермах.
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
	DeleteBiometrics(AnimalID int) error
	DeleteActivity(AnimalID int) error
}

// Farms описує методи, необхідні для операцій з фермами.
type Farms interface {
	Create(UserID int, farm farmsage.Farm) (int, error)
	GetAll(UserID int) ([]farmsage.Farm, error)
	GetByID(UserID, FarmID int) (farmsage.Farm, error)
	Delete(UserID, FarmID int) error
	Update(UserID, id int, input farmsage.UpdateFarmInput) error
}

// Feed описує методи, необхідні для управління кормами на фермах.
type Feed interface {
	Create(feed farmsage.Feed) (int, error)
	GetAll() ([]farmsage.Feed, error)
	Delete(feedID int) error
	Update(feedID int, input farmsage.UpdateFeedInput) error
}

// FeedingSchedule описує методи, пов'язані з графіком годування тварин на фермі.
type FeedingSchedule interface {
	Create(feedingSchedule farmsage.FeedingSchedule) (int, error)
	GetByID(animalID int) ([]farmsage.FeedingSchedule, error)
	Delete(scheduleID int) error
	Update(scheduleID int, input farmsage.UpdateFeedingScheduleInput) error
}

// Admin описує методи для управління користувачами та резервним копіюванням даних.
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

// Service об'єднує реалізації інтерфейсів для взаємодії з репозиторіями.
type Service struct {
	Authorization
	Animals
	Farms
	Feed
	FeedingSchedule
	Admin
}

// NewService створює новий сервіс на основі репозиторіїв.
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
