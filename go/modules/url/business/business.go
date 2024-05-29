package business

import (
	"context"
	"github.com/nghiatrann0502/demo-redis/go/modules/url/model"
)

type URLRepository interface {
	CreateUrlShorten(ctx context.Context, data *model.CreateUrl) error
	GetUrlShorten(ctx context.Context, shortUrl string) (*model.URLData, error)
	UpdateUrlById(ctx context.Context, data model.UrlUpdate) (*string, error)
}

type URLBusiness interface {
	CreateUrlShorten(ctx context.Context, longUrl string) (*string, error)
	GetUrlShorten(ctx context.Context, shortUrl string) (*model.URLData, error)
	UpdateById(ctx context.Context, input model.UrlUpdate) (*string, error)
}

type URLCache interface {
	GetUrlShorten(shortUrl string) (*model.URLData, error)
	GetUrlShortenV2(shortUrl string) (*model.URLData, error)
	CreateUrlShorten(data *model.URLData) error
	CreateUrlShortenV2(data *model.URLData) error
	DeleteCache(shortUrl string) error
}

type business struct {
	repo  URLRepository
	cache URLCache
}

func NewBusiness(repo URLRepository, cache URLCache) *business {
	return &business{
		repo:  repo,
		cache: cache,
	}
}
