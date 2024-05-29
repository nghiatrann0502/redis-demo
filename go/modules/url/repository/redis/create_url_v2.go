package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nghiatrann0502/demo-redis/go/modules/url/model"
)

func (storage *redisStorage) CreateUrlShortenV2(data *model.URLData) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return storage.db.Set(context.TODO(), fmt.Sprintf("URLSTRING:%s", data.ShortURL), body, 0).Err()
}
