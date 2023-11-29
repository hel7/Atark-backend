package service

import "github.com/hel7/Atark-backend/pkg/repository"

type Authorization interface {
	//	Register(username, password string) error
	//	Login(username, password string) (string, error)
	//	Logout(token string) error
}

type Animals interface {
	//	GetAllAnimals() ([]Animal, error)
	//	GetAnimalByID(id string) (Animal, error)
	//	CreateAnimal(details AnimalDetails) (string, error)
	//	UpdateAnimal(id string, details AnimalDetails) error
	//	DeleteAnimal(id string) error
}

type Farms interface {
	//	GetAllFarms() ([]Farm, error)
	//	GetFarmByID(id string) (Farm, error)
	//	CreateFarm(details FarmDetails) (string, error)
	//	UpdateFarm(id string, details FarmDetails) error
	//	DeleteFarm(id string) error
	//	AddAnimalToFarm(animalID, farmID string) error
	//	GetFarmForAnimal(animalID string) (string, error)
}

type Feed interface {
	//	GetAllFeed() ([]FeedItem, error)
	//	GetFeedByID(id string) (FeedItem, error)
	//	CreateFeed(details FeedDetails) (string, error)
	//	UpdateFeed(id string, details FeedDetails) error
	//	DeleteFeed(id string) error
}

type FeedingSchedule interface {
	//	GetAllFeedingSchedules() ([]FeedingScheduleItem, error)
	//	GetFeedingScheduleByID(id string) (FeedingScheduleItem, error)
	//	CreateFeedingSchedule(details FeedingScheduleDetails) (string, error)
	//	UpdateFeedingSchedule(id string, details FeedingScheduleDetails) error
	//	DeleteFeedingSchedule(id string) error
}

type Analytics interface {
	//	GetAnalytics() (AnalyticsData, error)
	//	GetAnalyticsByDate(date string) (AnalyticsData, error)
}

type Admin interface {
	//	GetAllUsers() ([]User, error)
	//	GetUserByID(id string) (User, error)
	//	CreateUser(details UserDetails) (string, error)
	//	UpdateUser(id string, details UserDetails) error
	//	DeleteUser(id string) error

	//	GetAllRoles() ([]Role, error)
	//	GetRoleByID(id string) (Role, error)
	//	CreateRole(details RoleDetails) (string, error)
	//	UpdateRole(id string, details RoleDetails) error
	//	DeleteRole(id string) error

	//	CreateBackup() error
	//	RestoreBackup() error
	//	ExportData() (string, error)
	//	ImportData(data string) error
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
	return &Service{}
}
