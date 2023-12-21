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

func (s *AnimalService) Create(FarmID int, animal farmsage.Animal) (int, error) {
	return s.repo.Create(FarmID, animal)
}

func (s *AnimalService) GetAll(FarmID int) ([]farmsage.Animal, error) {
	return s.repo.GetAll(FarmID)
}

func (s *AnimalService) GetByID(FarmID, AnimalID int) (farmsage.Animal, error) {
	return s.repo.GetByID(FarmID, AnimalID)
}
func (s *AnimalService) Delete(FarmID, AnimalID int) error {
	return s.repo.Delete(FarmID, AnimalID)
}
func (s *AnimalService) Update(AnimalID int, input farmsage.UpdateAnimalInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(AnimalID, input)
}
