package repository

type Authorization interface {
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

func NewRepository() *Repository {
	return &Repository{}
}
