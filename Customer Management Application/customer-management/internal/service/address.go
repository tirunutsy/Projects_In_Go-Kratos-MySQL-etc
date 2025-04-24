package service

import (
	"context"
	pb "customer-management/api/helloworld/v1/address"
	"customer-management/internal/biz"
)

type AddressService struct {
	pb.UnimplementedAddressServer
	biz *biz.Addressbiz
}

func NewAddressService(biz *biz.Addressbiz) *AddressService {
	return &AddressService{
		biz: biz,
	}
}

func (s *AddressService) CreateAddress(ctx context.Context, req *pb.CreateAddressRequest) (*pb.CreateAddressReply, error) {
	return s.biz.CreateAddress(ctx, req)
}
func (s *AddressService) UpdateAddress(ctx context.Context, req *pb.UpdateAddressRequest) (*pb.UpdateAddressReply, error) {
	return &pb.UpdateAddressReply{}, nil
}
func (s *AddressService) DeleteAddress(ctx context.Context, req *pb.DeleteAddressRequest) (*pb.DeleteAddressReply, error) {
	return &pb.DeleteAddressReply{}, nil
}
func (s *AddressService) GetAddress(ctx context.Context, req *pb.GetAddressRequest) (*pb.GetAddressReply, error) {
	return &pb.GetAddressReply{}, nil
}
func (s *AddressService) ListAddress(ctx context.Context, req *pb.ListAddressRequest) (*pb.ListAddressReply, error) {
	return &pb.ListAddressReply{}, nil
}
