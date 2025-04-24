package data

import (
	"customer-management/internal/conf"

	model "customer-management/models"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewCustomerRepository,
	NewEmailRepository,
	NewPhoneRepository,
	NewAddressRepository,
)


// Data .
type Data struct {
	DB *gorm.DB
}

// NewData
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	dsn := c.Database.Source
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	if err := db.AutoMigrate(&model.Customer{}, &model.Email{}, &model.Phone{}, &model.Address{}, &model.RelationTable{}); err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{DB: db}, cleanup, nil
}
