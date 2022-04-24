package model

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID uint64 `gorm:"primaryKey"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Task struct {
	Model

	Summary     string
	IsCompleted bool
}
