package mysql

import (
	"context"
	"github.com/nghiatrann0502/demo-redis/go/modules/url/model"
	"time"
)

func (storage *mySqlStorage) CreateUrlShorten(ctx context.Context, data *model.CreateUrl) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `INSERT INTO urls (short_url, original_url, created_at, updated_at)
		values (?, ?, ?, ?)`

	_, err := storage.db.ExecContext(ctx, stmt,
		data.ShortURL,
		data.URL,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}

	//_, err = insertResult.LastInsertId()
	//if err != nil {
	//	return err
	//}

	return nil
}
