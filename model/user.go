package model

import (
	"time"

	"github.com/wcl48/valval"
)

// User is userdata format
type User struct {
	ID          int64
	Name        string `sql:"size:255"`
	Description string `sql:"type:text"`
	CreateAt    time.Time
	UpdateAt    time.Time
	DeleteAt    time.Time
}

// UserValidate is
/*
 * userdata用バリテータ
 */
func UserValidate(user User) error {
	Validator := valval.Object(valval.M{
		"Name": valval.String(
			valval.MaxLength(50),
		),
		"description": valval.String(
			valval.MaxLength(10000),
		),
	})
	return Validator.Validate(user)
}
