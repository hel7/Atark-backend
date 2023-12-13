package service

import (
	farmsage "github.com/hel7/Atark-backend"
	"github.com/hel7/Atark-backend/pkg/repository"
)

type FeedService struct {
	repo repository.Feed
}

func NewFeedService(repo repository.Feed) *FeedService {
	return &FeedService{repo: repo}
}

func (s *FeedService) Create(feed farmsage.Feed) (int, error) {
	return s.repo.Create(feed)
}

func (s *FeedService) GetAll() ([]farmsage.Feed, error) {
	return s.repo.GetAll()
}
func (s *FeedService) Delete(FeedID int) error {
	return s.repo.Delete(FeedID)
}
func (s *FeedService) Update(feedID int, input farmsage.UpdateFeedInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(feedID, input)
}
