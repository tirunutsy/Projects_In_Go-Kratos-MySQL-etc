package biz

import (
	"context"
	pb "customer-management/api/helloworld/v1"
	"customer-management/internal/data"
	"customer-management/models"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	//"gorm.io/gorm"
)

type Customerbiz struct {
	customerdata *data.CustomerRepository
}

func NewCustomerBiz(customerdata *data.CustomerRepository) *Customerbiz {
	return &Customerbiz{
		customerdata: customerdata,
	}
}

func (b *Customerbiz) CreateCustomer(ctx context.Context, req *pb.CreateCustomerRequest) (*pb.CreateCustomerReply, error) {
	// Map the request to a Customer model
	customer := CustomerMapper(req)

	// Create the customer in the database
	err := b.customerdata.Create(customer)
	if err != nil {
		log.Printf("Failed to create customer: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to create customer: %v", err)
	}

	// Initialize var for related models
	var emailModels []models.Email
	var phoneModels []models.Phone
	var addressModels []models.Address
	var relationtable []models.RelationTable

	// Map and creating emails
	if req.Emails != nil {
		emailModels = convertProtoMailToModelMail(req.Emails)
		// Updating Email Table based on its Customer ID
		for i := range emailModels {
			emailModels[i].CustomerID = customer.ID
		}
		err = b.customerdata.CreateEmails(emailModels)
		if err != nil {
			log.Printf("Failed to create emails: %v", err)
			return nil, status.Errorf(codes.Internal, "failed to create emails: %v", err)
		}

		// Updating Relational Table details based on Relation type
		for _, email := range emailModels {
			relationtable = append(relationtable, models.RelationTable{
				CustomerID:   customer.ID,
				RelationType: "email",
				RelationID:   email.ID, // Use the email ID
			})
		}
	}

	// Map and creating phones
	if req.Phones != nil {
		phoneModels = convertProtoPhoneToModelPhone(req.Phones)
		// Updating phone table Numbers's Customer ID
		for i := range phoneModels {
			phoneModels[i].CustomerID = customer.ID
		}
		err = b.customerdata.CreatePhones(phoneModels)
		if err != nil {
			log.Printf("Failed to create phones: %v", err)
			return nil, status.Errorf(codes.Internal, "failed to create phones: %v", err)
		}

		// Updating Relational Table details based on Relation type
		for _, phone := range phoneModels {
			relationtable = append(relationtable, models.RelationTable{
				CustomerID:   customer.ID,
				RelationType: "phone",
				RelationID:   phone.ID,
			})
		}
	}

	// Map and create addresses
	if req.Addresses != nil {

		addressModels = convertProtoAddressToModelAddress(req.Addresses)

		// Updating Address Table based on its Customer ID
		for i := range addressModels {
			// Validating bassed on Enum values
			if addressModels[i].Type != "permanant" && addressModels[i].Type != "temperory" && addressModels[i].Type != "other" {
				return nil, status.Errorf(codes.InvalidArgument, "invalid address type: %s", addressModels[i].Type)
			}
			addressModels[i].CustomerID = customer.ID

		}
		err = b.customerdata.CreateAddresses(addressModels)
		if err != nil {
			log.Printf("Failed to create addresses: %v", err)
			return nil, status.Errorf(codes.Internal, "failed to create addresses: %v", err)
		}
		// Updating Relational Table details based on Relation type
		for _, add := range addressModels {
			relationtable = append(relationtable, models.RelationTable{
				CustomerID:   customer.ID,
				RelationType: "address",
				RelationID:   add.ID,
			})
		}
	}

	// Create relation table entries
	if len(relationtable) > 0 {
		err = b.customerdata.CreateRelations(relationtable)
		if err != nil {
			log.Printf("Failed to create relation table entries: %v", err)
			return nil, status.Errorf(codes.Internal, "failed to create relation table entries: %v", err)
		}
	}

	// for i := range relationtable {
	// 	switch relationtable[i].RelationType {
	// 	case "email":

	// 	}
	// }

	// Response
	reply := &pb.CustomerData{
		Id:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Age:       customer.Age,
		Gender:    customer.Gender,
		Emails:    convertModelMailToProtoMail(emailModels),
		Phones:    convertModelPhoneToProtoPhone(phoneModels),
		Addresses: convertModelAddressToProtoAddress(addressModels),
	}

	return &pb.CreateCustomerReply{
		Customer: reply,
	}, nil
}

func (b *Customerbiz) GetCustomer(ctx context.Context, req *pb.GetCustomerRequest) (*pb.GetCustomerReply, error) {
	log.Printf("Fetching customer with ID: %d", req.Id)

	// Get customer by id
	customer, err := b.customerdata.GetByID(req.Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("Customer with ID %d not found", req.Id)
			return nil, status.Errorf(codes.NotFound, "customer with ID %d not found", req.Id)
		}

		log.Printf("Error fetching customer: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to fetch customer: %v", err)
	}

	// get emails
	emails, err := b.customerdata.GetEmailsByCustomerID(customer.ID)
	if err != nil {
		log.Printf("Error fetching emails for customer ID %d: %v", customer.ID, err)
		return nil, status.Errorf(codes.Internal, "failed to fetch emails: %v", err)
	}

	// get phones
	phones, err := b.customerdata.GetPhonesByCustomerID(customer.ID)
	if err != nil {
		log.Printf("Error fetching phones for customer ID %d: %v", customer.ID, err)
		return nil, status.Errorf(codes.Internal, "failed to fetch phones: %v", err)
	}

	// get addresses
	addresses, err := b.customerdata.GetAddressesByCustomerID(customer.ID)
	if err != nil {
		log.Printf("Error fetching addresses for customer ID %d: %v", customer.ID, err)
		return nil, status.Errorf(codes.Internal, "failed to fetch addresses: %v", err)
	}

	// Convert those models to proto types
	emmmModel := convertModelMailToProtoMail(emails)
	phhhModel := convertModelPhoneToProtoPhone(phones)
	adddModel := convertModelAddressToProtoAddress(addresses)

	log.Printf("Customer fetched: %+v", customer)

	//  := &pb.CustomerData{
	// 	Id:        customer.ID,
	// 	FirstName: customer.FirstName,
	// 	LastName:  customer.LastName,
	// 	Age:       customer.Age,
	// 	Gender:    customer.Gender,

	reply := &pb.CustomerData{
		Id:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Age:       customer.Age,
		Gender:    customer.Gender,
		Emails:    emmmModel,
		Phones:    phhhModel,
		Addresses: adddModel,
	}

	return &pb.GetCustomerReply{
		Customer: reply,
	}, nil
}

func (b *Customerbiz) UpdateCustomer(ctx context.Context, req *pb.UpdateCustomerRequest) (*pb.UpdateCustomerReply, error) {
	customer := &models.Customer{
		ID:        req.Id,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Age:       req.Age,
		Gender:    req.Gender,
	}
	err := b.customerdata.Update(customer)
	if err != nil {
		return nil, nil
	}
	reply := &pb.CustomerData{
		Id:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Age:       customer.Age,
		Gender:    customer.Gender,
	}
	return &pb.UpdateCustomerReply{
		Customer: reply,
	}, nil
}

func (b *Customerbiz) DeleteCustomer(ctx context.Context, req *pb.DeleteCustomerRequest) (*pb.DeleteCustomerReply, error) {
	return &pb.DeleteCustomerReply{}, nil
}

func (b *Customerbiz) ListCustomer(ctx context.Context, req *pb.ListCustomersRequest) (*pb.ListCustomersReply, error) {
	return &pb.ListCustomersReply{}, nil
}

func convertProtoMailToModelMail(protomail []*pb.Emaildata) []models.Email {
	var modelmail []models.Email
	for _, value := range protomail {
		modelmail = append(modelmail, models.Email{
			// ID:        2232425,
			Email:     value.Email,
			Type:      value.Type,
			IsPrimary: value.IsPrimary,
		})
	}
	return modelmail
}

func convertModelMailToProtoMail(modelmail []models.Email) []*pb.Emaildata {
	var protomail []*pb.Emaildata
	for _, value := range modelmail {
		protomail = append(protomail, &pb.Emaildata{
			Id:        value.ID,
			Email:     value.Email,
			Type:      value.Type,
			IsPrimary: value.IsPrimary,
		})
	}
	return protomail
}

func convertProtoPhoneToModelPhone(protophone []*pb.Phones) []models.Phone {
	var modelphone []models.Phone
	for _, value := range protophone {
		modelphone = append(modelphone, models.Phone{
			// ID:        2232426,
			Phone:     value.Number,
			IsPrimary: value.IsPrimary,
			Type:      value.Type,
		})
	}
	return modelphone
}

func convertModelPhoneToProtoPhone(modelphone []models.Phone) []*pb.Phones {
	var protophone []*pb.Phones
	for _, value := range modelphone {
		protophone = append(protophone, &pb.Phones{
			Id:        value.ID,
			Number:    value.Phone,
			IsPrimary: value.IsPrimary,
			Type:      value.Type,
		})
	}
	return protophone
}

func convertProtoAddressToModelAddress(protoadd []*pb.Address) []models.Address {
	var modeladd []models.Address
	for _, value := range protoadd {
		modeladd = append(modeladd, models.Address{
			// ID:      2232427,
			Address: value.Address,
			Pincode: value.Pincode,
			Type:    value.Type,
		})
	}
	return modeladd
}

func convertModelAddressToProtoAddress(modeladd []models.Address) []*pb.Address {
	var protoadd []*pb.Address
	for _, value := range modeladd {
		protoadd = append(protoadd, &pb.Address{
			Id:      value.ID,
			Address: value.Address,
			Pincode: value.Pincode,
			Type:    value.Type,
		})
	}
	return protoadd
}
