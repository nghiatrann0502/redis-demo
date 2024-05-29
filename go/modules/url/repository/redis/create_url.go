package redis

import (
	"context"
	"fmt"
	"github.com/nghiatrann0502/demo-redis/go/modules/url/model"
)

func (storage *redisStorage) CreateUrlShorten(data *model.URLData) error {
	return storage.db.HSet(context.TODO(), fmt.Sprintf("URLS:%s", data.ShortURL), data).Err()
}
