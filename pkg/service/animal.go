package service

import (
	farmsage "github.com/hel7/Atark-backend"
	"github.com/hel7/Atark-backend/pkg/repository"
)

type AnimalService struct {
	repo repository.Animals
}

func NewAnimalService(repo repository.Animals) *AnimalService {
	return &AnimalService{repo: repo}
}

func (s *AnimalService) Create(UserID int, animal farmsage.Animal) (int, error) {
	return s.repo.Create(UserID, animal)
}

func (s *AnimalService) GetAll(UserID int) ([]farmsage.Animal, error) {
	return s.repo.GetAll(UserID)
}

func (s *AnimalService) GetByID(UserID, AnimalID int) (farmsage.Animal, error) {
	return s.repo.GetByID(UserID, AnimalID)
}
func (s *AnimalService) Delete(UserID, AnimalID int) error {
	return s.repo.Delete(UserID, AnimalID)
}
func (s *AnimalService) Update(AnimalID int, input farmsage.UpdateAnimalInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(AnimalID, input)
}
