package model

import (
	"errors"
	"strings"
	"time"
)

type Url struct {
	ID        int        `json:"id"`
	ShortURL  string     `json:"short_url"`
	URL       string     `json:"url"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type CreateUrl struct {
	ShortURL string `json:"short_url"`
	URL      string `json:"url"`
}

type UrlUpdate struct {
	ShortURL string `json:"short_url"`
	URL      string `json:"url"`
}

type URLData struct {
	ShortURL string `json:"short_url" redis:"short_url"`
	URL      string `json:"url" redis:"url"`
}

type UrlInput string

func (input *CreateUrl) Validate() error {
	if strings.TrimSpace(input.URL) == "" {
		return errors.New("url cannot be blank")
	}

	if strings.Split(input.URL, "://")[0] != "https" && strings.Split(input.URL, "://")[0] != "http" {
		return errors.New("url must start with http or https")
	}

	return nil
}

var NotFound = errors.New("data not found")
var RedisNotFound = errors.New("not found")
