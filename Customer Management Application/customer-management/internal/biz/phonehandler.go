package biz

import (
	"context"
	pb "customer-management/api/helloworld/v1/phone"
	"customer-management/internal/data"
	"customer-management/models"
)

type Phonebiz struct {
	phonedata *data.PhoneRepository
}

func NewPhonebiz(phonedata *data.PhoneRepository) *Phonebiz {
	return &Phonebiz{
		phonedata: phonedata,
	}
}
func (h *Phonebiz) CreatePhone(ctx context.Context, req *pb.CreatePhoneRequest) (*pb.CreatePhoneReply, error) {
	phone := &models.Phone{
		Phone: req.Phone,
	}
	err := h.phonedata.Create(phone)
	if err != nil {
		return nil, err
	}
	return &pb.CreatePhoneReply{
		Phone: phone.Phone,
	}, nil
}
func (h *Phonebiz) GetPhone(ctx context.Context, req *pb.GetPhoneRequest) (*pb.GetPhoneReply, error) {
	phone := &models.Phone{
		ID: req.Id,
	}
	phone, err := h.phonedata.Get(phone.Phone)
	if err != nil {
		return nil, err
	}
	return &pb.GetPhoneReply{
		Phone: phone.Phone,
	}, nil
}

func (h *Phonebiz) UpdatePhone(ctx context.Context, req *pb.UpdatePhoneRequest) (*pb.UpdatePhoneReply, error) {
	phone := &models.Phone{
		ID: req.Id,
	}
	err := h.phonedata.Update(phone)
	if err != nil {
		return nil, err
	}
	return &pb.UpdatePhoneReply{
		Phone: phone.Phone,
	}, nil
}

func (h *Phonebiz) DeletePhone(ctx context.Context, req *pb.DeletePhoneRequest) (*pb.DeletePhoneReply, error) {

	phone := &models.Phone{
		ID: req.Id,
	}

	err := h.phonedata.Delete(phone.ID)
	if err != nil {
		return nil, err
	}
	message := "Phone deleted successfully"
	return &pb.DeletePhoneReply{
		Message: message,
	}, nil
}

func (h *Phonebiz) ListPhone(ctx context.Context, req *pb.ListPhoneRequest) (*pb.ListPhoneReply, error) {
	phones, err := h.phonedata.List()
	if err != nil {
		return nil, err
	}
	var phoneList []*pb.PhoneList
	for _, phone := range phones {
		phoneList = append(phoneList, &pb.PhoneList{
			Id:    phone.ID,
			Phone: phone.Phone,
		})
	}
	return &pb.ListPhoneReply{
		Listphones: phoneList,
	}, nil
}
