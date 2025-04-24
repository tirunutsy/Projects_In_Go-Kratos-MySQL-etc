package data

import (
	model "customer-management/models"
)

type PhoneRepository struct {
	db *Data
}
func NewPhoneRepository(db *Data) *PhoneRepository {	
	return &PhoneRepository{
		db: db,
	}
}
func (r *PhoneRepository) Create(phone *model.Phone) error {
	err := r.db.DB.Create(phone).Error
	if err != nil {
		return err
	}
	return nil
}

func(r *PhoneRepository) Delete(id uint32) error {
	var phone model.Phone
	err := r.db.DB.First(&phone, id).Error
	if err != nil {
		return err
	}
	err = r.db.DB.Delete(&phone).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *PhoneRepository) Get(id string) (*model.Phone, error) {
	var phone model.Phone
	err := r.db.DB.First(&phone, id).Error
	if err != nil {
		return nil, err
	}
	return &phone, nil
}
func (r *PhoneRepository) Update(phone *model.Phone) error {
	err := r.db.DB.Save(phone).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *PhoneRepository) List() ([]model.Phone, error) {
	var phones []model.Phone
	err := r.db.DB.Find(&phones).Error
	if err != nil {
		return nil, err
	}
	return phones, nil
}
