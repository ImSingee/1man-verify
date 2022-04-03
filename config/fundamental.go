package config

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var (
	Validator *validator.Validate
	DB        *gorm.DB
)

func init() {
	Validator = validator.New()
}
