package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nghiatrann0502/demo-redis/go/modules/url/model"
	"github.com/redis/go-redis/v9"
)

func (storage *redisStorage) GetUrlShortenV2(shortUrl string) (*model.URLData, error) {
	var data model.URLData
	result, err := storage.db.
		Get(context.TODO(), fmt.Sprintf("URLSTRING:%s", shortUrl)).
		Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, model.RedisNotFound
		}
		return nil, err
	}

	if err := json.Unmarshal([]byte(result), &data); err != nil {
		return nil, err
	}

	if data.URL == "" {
		return nil, model.RedisNotFound
	}

	return &data, nil
}
