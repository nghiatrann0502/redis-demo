package business

import (
	"context"
	"github.com/nghiatrann0502/demo-redis/go/common"
	"github.com/nghiatrann0502/demo-redis/go/modules/url/model"
)

func (biz *business) CreateUrlShorten(ctx context.Context, longUrl string) (*string, error) {
	shortUrl := common.GenerateShortLink(longUrl)
	data := &model.CreateUrl{URL: longUrl, ShortURL: shortUrl}

	if err := data.Validate(); err != nil {
		return nil, err
	}

	if err := biz.repo.CreateUrlShorten(ctx, data); err != nil {
		return nil, err
	}

	return &shortUrl, nil
}
