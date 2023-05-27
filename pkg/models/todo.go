package models

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Name        string `json:"name" validate:"required" gorm:"not null"`
	Description string `json:"description" validate:"required"`
	Completed   bool   `json:"completed"`
}

func (todo *Todo) Validate() error {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	return validate.Struct(todo)
}

func (todo *Todo) Create(db *gorm.DB) error {
	return db.Create(todo).Error
}
