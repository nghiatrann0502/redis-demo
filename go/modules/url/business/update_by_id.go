package business

import (
	"context"
	"github.com/nghiatrann0502/demo-redis/go/modules/url/model"
)

func (biz *business) UpdateById(ctx context.Context, input model.UrlUpdate) (*string, error) {
	_, err := biz.repo.GetUrlShorten(ctx, input.ShortURL)

	if err != nil {
		return nil, err
	}

	if err := biz.cache.DeleteCache(input.ShortURL); err != nil {
		return nil, err
	}

	shortUrl, err := biz.repo.UpdateUrlById(ctx, input)
	if err != nil {
		return nil, err
	}

	return shortUrl, nil
}
