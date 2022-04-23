package model

import "gorm.io/gorm"

var (
	invalidDeletedAt           = gorm.DeletedAt{Valid: false}
	invalidDeletedAtAsMapValue = nullTimeAsMapValue(invalidDeletedAt)
)
