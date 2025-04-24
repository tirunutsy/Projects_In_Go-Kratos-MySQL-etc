package service

import (
	"context"

	pb "customer-management/api/helloworld/v1/phone"
	"customer-management/internal/biz"
)

type PhoneService struct {
	pb.UnimplementedPhoneServiceServer
	biz *biz.Phonebiz
}

func NewPhoneService(biz *biz.Phonebiz) *PhoneService {
	return &PhoneService{
		biz: biz,
	}
}

func (s *PhoneService) CreatePhone(ctx context.Context, req *pb.CreatePhoneRequest) (*pb.CreatePhoneReply, error) {
	return s.biz.CreatePhone(ctx, req)
}
func (s *PhoneService) UpdatePhone(ctx context.Context, req *pb.UpdatePhoneRequest) (*pb.UpdatePhoneReply, error) {
	return &pb.UpdatePhoneReply{}, nil
}
func (s *PhoneService) DeletePhone(ctx context.Context, req *pb.DeletePhoneRequest) (*pb.DeletePhoneReply, error) {
	return &pb.DeletePhoneReply{}, nil
}
func (s *PhoneService) GetPhone(ctx context.Context, req *pb.GetPhoneRequest) (*pb.GetPhoneReply, error) {
	return &pb.GetPhoneReply{}, nil
}
func (s *PhoneService) ListPhone(ctx context.Context, req *pb.ListPhoneRequest) (*pb.ListPhoneReply, error) {
	return &pb.ListPhoneReply{}, nil
}
