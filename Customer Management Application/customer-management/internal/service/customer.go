package service

import (
	"context"
	pb "customer-management/api/helloworld/v1"
	"customer-management/internal/biz"
)

type CustomerService struct {
	pb.UnimplementedCustomerServer
	biz *biz.Customerbiz
}

func NewCustomerService(biz *biz.Customerbiz) *CustomerService {
	return &CustomerService{biz: biz}
}

func (s *CustomerService) CreateCustomer(ctx context.Context, req *pb.CreateCustomerRequest) (*pb.CreateCustomerReply, error) {
	return s.biz.CreateCustomer(ctx, req)
}
func (s *CustomerService) UpdateCustomer(ctx context.Context, req *pb.UpdateCustomerRequest) (*pb.UpdateCustomerReply, error) {
	return &pb.UpdateCustomerReply{}, nil
}
func (s *CustomerService) DeleteCustomer(ctx context.Context, req *pb.DeleteCustomerRequest) (*pb.DeleteCustomerReply, error) {
	return &pb.DeleteCustomerReply{}, nil
}
func (s *CustomerService) GetCustomer(ctx context.Context, req *pb.GetCustomerRequest) (*pb.GetCustomerReply, error) {
	return s.biz.GetCustomer(ctx, req)
}
func (s *CustomerService) ListCustomer(ctx context.Context, req *pb.ListCustomersRequest) (*pb.ListCustomersReply, error) {
	return &pb.ListCustomersReply{}, nil
}
