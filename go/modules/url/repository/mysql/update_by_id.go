package mysql

import (
	"context"
	"github.com/nghiatrann0502/demo-redis/go/modules/url/model"
	"log"
)

func (storage *mySqlStorage) UpdateUrlById(ctx context.Context, data model.UrlUpdate) (*string, error) {
	query := "UPDATE urls SET original_url=? where short_url=?"
	res, err := storage.db.ExecContext(ctx, query, data.URL, data.ShortURL)
	if err != nil {
		return nil, err
	}
	count, err := res.RowsAffected()

	if err != nil {
		log.Println(err)
	}

	log.Println(count)
	return nil, nil
}
