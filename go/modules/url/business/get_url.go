package business

import (
	"context"
	"errors"
	"github.com/nghiatrann0502/demo-redis/go/modules/url/model"
	"log"
)

func (biz *business) GetUrlShorten(ctx context.Context, shortUrl string) (*model.URLData, error) {

	data, err := biz.cache.GetUrlShorten(shortUrl)
	if err != nil && !errors.Is(err, model.RedisNotFound) {
		return nil, err
	}

	if data != nil {
		return data, nil
	}

	data, err = biz.repo.GetUrlShorten(ctx, shortUrl)
	if err != nil {
		return nil, err
	}

	if err := biz.cache.CreateUrlShorten(data); err != nil {
		log.Println(err, "error when create cache")
	}

	return data, nil
}
