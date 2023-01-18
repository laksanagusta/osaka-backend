package organization

import "gorm.io/gorm"

type Repository interface {
	Save(organization Organization) (Organization, error)
	FindById(organizationId int) (Organization, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(organization Organization) (Organization, error) {
	err := r.db.Create(&organization).Error

	if err != nil {
		return organization, err
	}

	return organization, nil
}

func (r *repository) FindById(organizationId int) (Organization, error) {
	organization := Organization{}
	err := r.db.Preload("User").Where("id = ?", organizationId).Find(&organization).Error

	if err != nil {
		return organization, err
	}

	return organization, nil
}
