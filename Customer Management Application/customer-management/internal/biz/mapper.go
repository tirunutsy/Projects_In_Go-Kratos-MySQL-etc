package biz

import (
	pb "customer-management/api/helloworld/v1"
	"customer-management/models"
)

func CustomerMapper(req *pb.CreateCustomerRequest) *models.Customer {

	return &models.Customer{
		// ID:        2232232424,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Age:       req.Age,
		Gender:    req.Gender,
	}
}

func EmailMapper(req *pb.Emaildata) *models.Email {

	return &models.Email{
		Email: req.Email,
	}

	// Emails:    convertProtoMailToModelMail(req.Emails),
}

func PhoneMapper(req *pb.Phones) *models.Phone {
	// Phones:    convertProtoPhoneToModelPhone(req.Phones),

	return &models.Phone{Phone: req.Number}
}

func AddressMapper(req *pb.Address) *models.Address {
	return &models.Address{
		Address: req.Address,
		Pincode: req.Pincode,
		Type:    req.Type, // Map the 'type' field correctly
	}
}
