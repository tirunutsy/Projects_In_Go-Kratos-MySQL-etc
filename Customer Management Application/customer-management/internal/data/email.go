package data

import (
	model "customer-management/models"
)

type EmailRepository struct {
	db *Data
}

func NewEmailRepository(db *Data) *EmailRepository {
	return &EmailRepository{
		db: db,
	}
}
func (r *EmailRepository) Create(email *model.Email) error {
	err := r.db.DB.Create(&email).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *EmailRepository) GetAll() ([]model.Email, error) {
	var emails []model.Email
	err := r.db.DB.Find(&emails).Error
	if err != nil {
		return nil, err
	}
	return emails, nil
}
func (r *EmailRepository) GetByID(id string) (*model.Email, error) {
	var email model.Email
	err := r.db.DB.First(&email, id).Error
	if err != nil {
		return nil, err
	}
	return &email, nil
}
func (r *EmailRepository) Update(email *model.Email) error {
	err := r.db.DB.Save(email).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *EmailRepository) Delete(id uint32) error {
	var email model.Email
	err := r.db.DB.First(&email, id).Error
	if err != nil {
		return err
	}
	err = r.db.DB.Delete(&email).Error
	if err != nil {
		return err
	}
	return nil
}
