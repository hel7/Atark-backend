package service

import (
	farmsage "github.com/hel7/Atark-backend"
	"github.com/hel7/Atark-backend/pkg/repository"
)

type FarmService struct {
	repo repository.Farms
}

func NewFarmService(repo repository.Farms) *FarmService {
	return &FarmService{repo: repo}
}

func (s *FarmService) Create(UserID int, farm farmsage.Farm) (int, error) {
	return s.repo.Create(UserID, farm)
}

func (s *FarmService) GetAll(UserID int) ([]farmsage.Farm, error) {
	return s.repo.GetAll(UserID)
}

func (s *FarmService) GetByID(UserID, FarmID int) (farmsage.Farm, error) {
	return s.repo.GetByID(UserID, FarmID)
}
func (s *FarmService) Delete(UserID, FarmID int) error {
	return s.repo.Delete(UserID, FarmID)
}
