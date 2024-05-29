package redis

import (
	"context"
	"fmt"
)

func (storage *redisStorage) DeleteCache(shortUrl string) error {
	return storage.db.Del(context.Background(), fmt.Sprintf("URLS:%s", shortUrl)).Err()
}
