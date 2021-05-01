package models

import (
	"gorm.io/gorm"
)

// gorm.Model definition
type Order struct {
	gorm.Model
	ID                  uint
	Asunto              string
	DescripcionProblema string
	ClientID            uint
}
