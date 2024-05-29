package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/nghiatrann0502/demo-redis/go/modules/url/model"
	"github.com/redis/go-redis/v9"
)

func (storage *redisStorage) GetUrlShorten(shortUrl string) (*model.URLData, error) {
	var data model.URLData
	err := storage.db.
		HMGet(context.TODO(), fmt.Sprintf("URLS:%s", shortUrl), "short_url", "url").
		Scan(&data)

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, model.RedisNotFound
		}
		return nil, err
	}

	if data.URL == "" {
		return nil, model.RedisNotFound
	}

	return &data, nil
}
