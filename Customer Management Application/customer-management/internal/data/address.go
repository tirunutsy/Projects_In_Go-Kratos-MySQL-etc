package data

import (
	model "customer-management/models"
)

type AddressRepository struct {
	db *Data
}

func NewAddressRepository(db *Data) *AddressRepository {
	return &AddressRepository{
		db: db,
	}
}
func (r *AddressRepository) Create(address *model.Address) error {
	err := r.db.DB.Create(address).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *AddressRepository) GetByID(id uint32) (*model.Address, error) {
	var address model.Address
	err := r.db.DB.First(&address, id).Error
	if err != nil {
		return nil, err
	}
	return &address, nil
}

func (r *AddressRepository) Update(address *model.Address) error {
	err := r.db.DB.Save(address).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *AddressRepository) Delete(id uint32) error {
	var address model.Address
	err := r.db.DB.First(&address, id).Error
	if err != nil {
		return err
	}
	err = r.db.DB.Delete(&address).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *AddressRepository) List() ([]model.Address, error) {
	var addresses []model.Address
	err := r.db.DB.Find(&addresses).Error
	if err != nil {
		return nil, err
	}
	return addresses, nil
}
