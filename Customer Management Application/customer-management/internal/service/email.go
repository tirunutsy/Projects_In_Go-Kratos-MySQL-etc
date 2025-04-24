package service

import (
	"context"
	pb "customer-management/api/helloworld/v1"
	"customer-management/internal/biz"
)

type EmailService struct {
	pb.UnimplementedEmailServer
	biz *biz.Emailbiz
}

func NewEmailService(biz *biz.Emailbiz) *EmailService {
	return &EmailService{
		biz: biz,
	}
}

func (s *EmailService) CreateEmail(ctx context.Context, req *pb.CreateEmailRequest) (*pb.CreateEmailReply, error) {
	return s.biz.CreateEmail(ctx, req)
}
func (s *EmailService) UpdateEmail(ctx context.Context, req *pb.UpdateEmailRequest) (*pb.UpdateEmailReply, error) {
	return &pb.UpdateEmailReply{}, nil
}
func (s *EmailService) DeleteEmail(ctx context.Context, req *pb.DeleteEmailRequest) (*pb.DeleteEmailReply, error) {
	return &pb.DeleteEmailReply{}, nil
}
func (s *EmailService) GetEmail(ctx context.Context, req *pb.GetEmailRequest) (*pb.GetEmailReply, error) {
	return &pb.GetEmailReply{}, nil
}
func (s *EmailService) ListEmail(ctx context.Context, req *pb.ListEmailRequest) (*pb.ListEmailReply, error) {
	return &pb.ListEmailReply{}, nil
}
