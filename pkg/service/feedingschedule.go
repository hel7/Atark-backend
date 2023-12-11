package service

import (
	farmsage "github.com/hel7/Atark-backend"
	"github.com/hel7/Atark-backend/pkg/repository"
)

type FeedingScheduleService struct {
	repo repository.FeedingSchedule
}

func NewFeedingScheduleService(repo repository.FeedingSchedule) *FeedingScheduleService {
	return &FeedingScheduleService{repo: repo}
}

func (s *FeedingScheduleService) Create(feedingSchedule farmsage.FeedingSchedule) (int, error) {
	return s.repo.Create(feedingSchedule)
}

func (s *FeedingScheduleService) GetByID(animalID int) ([]farmsage.FeedingSchedule, error) {
	return s.repo.GetByID(animalID)
}
func (s *FeedingScheduleService) Delete(scheduleID int) error {
	return s.repo.Delete(scheduleID)
}
