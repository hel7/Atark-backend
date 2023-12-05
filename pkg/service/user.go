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
