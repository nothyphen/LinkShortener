package services

import (
	"linkshortner/repository"
	"linkshortner/serilizers"
)

type ShortService interface {
	ShowLink(request serilizers.RedirectRequest) (string, error)
}

type shortService struct {
	shortRepository repository.LinkRepository
}

func NewShortRepository(Repository repository.LinkRepository) ShortService {
	return &shortService{
		shortRepository: Repository,
	}
}

func (c *shortService) ShowLink(request serilizers.RedirectRequest) (string, error) {
	linkExists, err := c.shortRepository.ShowLink(request.Short)
	if err != nil {
		return "link not found", nil
	}

	return linkExists, nil
}