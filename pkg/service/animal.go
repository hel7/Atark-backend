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

func (s *AnimalService) AddActivity(AnimalID int, activity farmsage.Activity) (int, error) {
	return s.repo.AddActivity(AnimalID, activity)
}

func (s *AnimalService) GetActivityByAnimalID(AnimalID int) ([]farmsage.Activity, error) {
	return s.repo.GetActivityByAnimalID(AnimalID)
}

func (s *AnimalService) AddBiometrics(AnimalID int, biometrics farmsage.Biometrics) (int, error) {
	return s.repo.AddBiometrics(AnimalID, biometrics)
}

func (s *AnimalService) GetBiometricsByAnimalID(AnimalID int) ([]farmsage.Biometrics, error) {
	return s.repo.GetBiometricsByAnimalID(AnimalID)
}

func (s *AnimalService) DeleteBiometrics(BiometricID int) error {
	return s.repo.DeleteBiometrics(BiometricID)
}

func (s *AnimalService) DeleteActivity(ActivityID int) error {
	return s.repo.DeleteActivity(ActivityID)
}
