package converter

import "gorm.io/gorm"

var (
	invalidNullTime         = gorm.DeletedAt{Valid: false}
	invalidNullTimeAsString = nullTimeToString(invalidNullTime)
)
