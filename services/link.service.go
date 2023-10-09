package services

import (
	"linkshortner/entity"
	"linkshortner/repository"
	"linkshortner/serilizers"
	"math/rand"
)

type LinkService interface {
	AddLinkService(request serilizers.ShortRequest) (string, error)
}

type linkService struct {
	linkRepository repository.LinkRepository
}

func NewLinkService(Repository repository.LinkRepository) LinkService {
	return &linkService{
		linkRepository: Repository,
	}
}

var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// n is the length of random string we want to generate
func randStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		// randomly select 1 character from given charset
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func (c *linkService) AddLinkService(request serilizers.ShortRequest) (string, error) {
	link := entity.Short{}
	linkExists, _ := c.linkRepository.FindByLink(request.Link)
	if linkExists != "" {
		return linkExists, nil
	}

	link.Link = request.Link
	link.Short = randStr(7)

	_ , err := c.linkRepository.AddLink(link)
	if err != nil {
		return "", err
	} 

	return link.Short, nil
}