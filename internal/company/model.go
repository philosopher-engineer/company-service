package company

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Company struct {
	ID                uuid.UUID `gorm:"type:uuid;primary_key"`
	Name              string    `gorm:"type:varchar(15);unique;not null"`
	Description       string    `gorm:"type:varchar(3000)"`
	AmountOfEmployees int       `gorm:"not null"`
	Registered        bool      `gorm:"not null"`
	Type              string    `gorm:"type:varchar(50);not null"`
}

func (company *Company) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.New())
}
