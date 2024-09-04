package helpers

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ValidateRecordByID(tableName, id, nullStr string, db *pgxpool.Pool) error {
	// database - de request gelen id bilen gabat gelyan maglumat barmy ya-da yokmy sol barlanyar
	// eger yok bolsa onda error return edilyar
	query := fmt.Sprintf("SELECT id FROM %s WHERE id = '%s' AND deleted_at IS %s", tableName, id, nullStr)
	if err := db.QueryRow(context.Background(), query).Scan(&id); err != nil {
		return errors.New("record not found")
	}
	return nil
}

func ValidateStructData(s interface{}) error {
	validate := validator.New()
	if err := validate.Struct(s); err != nil {
		return err
	}
	return nil
}

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
