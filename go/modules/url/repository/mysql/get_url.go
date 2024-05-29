package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/nghiatrann0502/demo-redis/go/modules/url/model"
)

func (storage *mySqlStorage) GetUrlShorten(ctx context.Context, shortUrl string) (*model.URLData, error) {
	var data model.URLData
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `SELECT short_url, original_url FROM urls WHERE short_url = ?`
	err := storage.db.QueryRowContext(ctx, stmt, shortUrl).Scan(&data.ShortURL, &data.URL)
	if err != nil {
		fmt.Println(err, "err")

		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.NotFound
		}

		return nil, err
	}

	return &data, nil
}
