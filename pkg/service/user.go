package service

import (
	farmsage "github.com/hel7/Atark-backend"
	"github.com/hel7/Atark-backend/pkg/repository"
)

type AdminService struct {
	repo repository.Admin
}

func NewAdminService(repo repository.Admin) *AdminService {
	return &AdminService{repo: repo}
}

func (s *AdminService) CreateUser(user farmsage.User) (int, error) {
	return s.repo.CreateUser(user)
}

func (s *AdminService) GetUserByID(userID int) (farmsage.User, error) {
	return s.repo.GetUserByID(userID)
}

func (s *AdminService) GetAllUsers() ([]farmsage.User, error) {
	return s.repo.GetAllUsers()
}
func (s *AdminService) Delete(UserID int) error {
	return s.repo.Delete(UserID)
}
func (s *AdminService) UpdateUser(UserID int, input farmsage.UpdateUserInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.UpdateUser(UserID, input)
}
func (s *AdminService) BackupData(backupPath string) error {
	err := s.repo.BackupData(backupPath)
	if err != nil {
		return err
	}

	return nil
}
func (s *AdminService) RestoreData(backupPath string) error {
	err := s.repo.RestoreData(backupPath)
	if err != nil {
		return err
	}

	return nil
}
func (s *AdminService) ExportData(exportPath string) error {
	err := s.repo.ExportData(exportPath)
	if err != nil {
		return err
	}

	return nil
}
func (s *AdminService) ImportData(importPath string) error {
	err := s.repo.ImportData(importPath)
	if err != nil {
		return err
	}

	return nil
}
