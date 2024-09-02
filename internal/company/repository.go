package company

import (
	"github.com/jinzhu/gorm"
)

type Repository interface {
	CreateCompany(company *Company) error
	GetCompanyByID(id string) (*Company, error)
	UpdateCompany(company *Company) error
	DeleteCompany(id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) CreateCompany(company *Company) error {
	return r.db.Create(company).Error
}

func (r *repository) GetCompanyByID(id string) (*Company, error) {
	var company Company
	err := r.db.Where("id = ?", id).First(&company).Error
	return &company, err
}

func (r *repository) UpdateCompany(company *Company) error {
	return r.db.Save(company).Error
}

func (r *repository) DeleteCompany(id string) error {
	return r.db.Where("id = ?", id).Delete(&Company{}).Error
}
