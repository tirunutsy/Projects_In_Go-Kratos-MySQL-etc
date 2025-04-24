package biz

import (
	"context"
	pb "customer-management/api/helloworld/v1"
	"customer-management/internal/data"
	"customer-management/models"
)

type Emailbiz struct {
	emaildata *data.EmailRepository
}

func NewEmailbiz(emaildata *data.EmailRepository) *Emailbiz {
	return &Emailbiz{
		emaildata: emaildata,
	}
}
func (h *Emailbiz) CreateEmail(ctx context.Context, req *pb.CreateEmailRequest) (*pb.CreateEmailReply, error) {
	email := &models.Email{
		Email: req.Email,
	}
	err := h.emaildata.Create(email)
	if err != nil {
		return nil, err
	}
	return &pb.CreateEmailReply{
		Email: email.Email,
	}, nil
}
func (h *Emailbiz) GetEmail(ctx context.Context, req *pb.GetEmailRequest) (*pb.GetEmailReply, error) {
	email := &models.Email{
		ID: req.Id,
	}
	email, err := h.emaildata.GetByID(email.Email)
	if err != nil {
		return nil, err
	}
	return &pb.GetEmailReply{
		Email: email.Email,
	}, nil
}

func (h *Emailbiz) ListEmail(ctx context.Context, req *pb.ListEmailRequest) (*pb.ListEmailReply, error) {
	emails, err := h.emaildata.GetAll()
	if err != nil {
		return nil, err
	}
	var emailList []*pb.Emails
	for _, email := range emails {
		emailList = append(emailList, &pb.Emails{
			Id:        uint32(email.ID),
			Email:     email.Email,
			Type:      email.Type,
			IsPrimary: email.IsPrimary,
			Userid:    uint32(email.CustomerID),
		})
	}
	return &pb.ListEmailReply{
		Emails: emailList,
	}, nil
}
func (h *Emailbiz) UpdateEmail(ctx context.Context, req *pb.UpdateEmailRequest) (*pb.UpdateEmailReply, error) {
	email := &models.Email{
		ID:    req.Id,
		Email: req.Email,
	}
	err := h.emaildata.Update(email)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateEmailReply{
		Email: email.Email,
	}, nil
}
func (h *Emailbiz) DeleteEmail(ctx context.Context, req *pb.DeleteEmailRequest) (*pb.DeleteEmailReply, error) {
	email := &models.Email{
		ID: req.Id,
	}
	err := h.emaildata.Delete(uint32(email.CustomerID))
	if err != nil {
		return nil, err
	}
	return &pb.DeleteEmailReply{
		Message: "Email deleted successfully",
	}, nil
}
