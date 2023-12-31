package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
)

type ShortUrl struct {
	ID        int64     `json:"id"`
	Url       string    `validate:"required,url" json:"url"`
	Code      string    `validate:"required,len=6" json:"code"`
	CreatedAt time.Time `json:"createdAt"`
}

type ShortUrlModel struct {
	DB *sql.DB
}

func (u *ShortUrl) IsValid() (bool, []ErrorResponse) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	var errors []ErrorResponse

	errs := validate.Struct(u)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			errors = append(errors, ErrorResponse{
				Field: err.Field(),
				Tag:   err.Tag(),
				Value: err.Value(),
			})

		}
		return false, errors
	}
	return true, errors
}

func (m *ShortUrlModel) Insert(u *ShortUrl) error {
	query := `
		INSERT INTO shorturl (url, code) 
		VALUES ($1, $2) 
		RETURNING id, created_at 
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, u.Url, u.Code).Scan(&u.ID, &u.CreatedAt)
}

func (m *ShortUrlModel) GetByCode(c string) (*ShortUrl, error) {
	query := `SELECT * FROM shorturl WHERE code = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var shortUrl ShortUrl
	err := m.DB.QueryRowContext(ctx, query, c).Scan(
		&shortUrl.ID,
		&shortUrl.Url,
		&shortUrl.Code,
		&shortUrl.CreatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("record not found")
		default:
			return nil, err
		}
	}

	return &shortUrl, nil
}
