package data

import (
	"customer-management/models"
	// model "customer-management/models"
	// model "customer-management/models"
	// "errors"
	// "gorm.io/gorm"
)

type CustomerRepository struct {
	db *Data
}

func NewCustomerRepository(db *Data) *CustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}
func (r *CustomerRepository) Create(customer *models.Customer) error {
	return r.db.DB.Create(customer).Error
}

func (r *CustomerRepository) GetAll() ([]models.Customer, error) {
	var customers []models.Customer
	err := r.db.DB.Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}
func (r *CustomerRepository) GetByID(id uint32) (*models.Customer, error) {
	var cust models.Customer
	result := r.db.DB.Where("id = ?", id).First(&cust).Error
	// result := r.db.DB.First(&cust, id)
	// if result.Error != nil {
	// 	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	// 		return nil, result.Error
	// 	}
	// 	return nil, result.Error
	// }
	if result != nil {
		return nil, result
	}
	return &cust, nil
}

func (r *CustomerRepository) Update(customer *models.Customer) error {
	err := r.db.DB.Save(customer).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *CustomerRepository) Delete(id uint32) error {
	var customer models.Customer
	err := r.db.DB.First(&customer, id).Error
	if err != nil {
		return err
	}
	err = r.db.DB.Delete(&customer).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *CustomerRepository) CreateEmails(emails []models.Email) error {
	return r.db.DB.Create(&emails).Error
}

func (r *CustomerRepository) CreatePhones(phones []models.Phone) error {
	return r.db.DB.Create(&phones).Error
}

func (r *CustomerRepository) CreateAddresses(addresses []models.Address) error {
	return r.db.DB.Create(&addresses).Error
}

func (r *CustomerRepository) CreateRelations(relations []models.RelationTable) error {
	return r.db.DB.Create(&relations).Error
}

func (r *CustomerRepository) GetEmailsByCustomerID(customerID uint32) ([]models.Email, error) {
	var emails []models.Email
	err := r.db.DB.Where("customer_id = ?", customerID).Find(&emails).Error
	if err != nil {
		return nil, err
	}
	return emails, nil
}

func (r *CustomerRepository) GetPhonesByCustomerID(customerID uint32) ([]models.Phone, error) {
	var phones []models.Phone
	err := r.db.DB.Where("customer_id = ?", customerID).Find(&phones).Error
	if err != nil {
		return nil, err
	}
	return phones, nil
}

func (r *CustomerRepository) GetAddressesByCustomerID(customerID uint32) ([]models.Address, error) {
	var addresses []models.Address
	err := r.db.DB.Where("customer_id = ?", customerID).Find(&addresses).Error
	if err != nil {
		return nil, err
	}
	return addresses, nil
}
