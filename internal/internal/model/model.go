package model

import (
	"time"

	"gorm.io/gorm"
)

type M map[string]interface{}

type Model struct {
	ID uint64 `gorm:"primaryKey"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func timeAsMapValue(input time.Time) string {
	return input.Format(time.RFC3339)
}

func nullTimeAsMapValue(input gorm.DeletedAt) *string {
	if !input.Valid {
		return nil
	}

	inputAsString := input.Time.Format(time.RFC3339)

	return &inputAsString
}
