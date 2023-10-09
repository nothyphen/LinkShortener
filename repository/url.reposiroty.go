package repository

import (
	"linkshortner/entity"

	"gorm.io/gorm"
)

type LinkRepository interface {
	AddLink(link entity.Short) (string, error)
	FindByLink(link string) (string, error)
	ShowLink(short string) (string, error)
}

type linkRepository struct {
	conn *gorm.DB
}

func NewLinkRepository(connection *gorm.DB) LinkRepository {
	return &linkRepository{
		conn: connection,
	}
}

func (c *linkRepository) AddLink(link entity.Short) (string, error) {
	if link.Link == "" {
		return "error : link is not found", nil
	}
	c.conn.Save(&link)
	return "inserted", nil
}

func (c *linkRepository) FindByLink(link string) (string, error) {
	var data entity.Short
	result := c.conn.Where("link = ?", link).Take(&data)

	if result.Error != nil {
		return "", result.Error
	}
	return data.Short, nil
}

func (c *linkRepository) ShowLink(short string) (string, error) {
	var data entity.Short

	result := c.conn.Where("short = ?", short).Take(&data)

	if result.Error != nil {
		return "", nil
	}
	return data.Link, nil
}