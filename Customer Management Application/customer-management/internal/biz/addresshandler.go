package biz

import (
	"context"
	pb "customer-management/api/helloworld/v1/address"
	"customer-management/internal/data"
	"customer-management/models"
)

type Addressbiz struct {
	addressdata *data.AddressRepository
}

func NewAddressBiz(addressdata *data.AddressRepository) *Addressbiz {
	return &Addressbiz{
		addressdata: addressdata,
	}
}
func (b *Addressbiz) CreateAddress(ctx context.Context, req *pb.CreateAddressRequest) (*pb.CreateAddressReply, error) {
	address := &models.Address{
		Address: req.Address,
		Pincode: req.Pincode,
		Type:    req.Type,
	}
	err := b.addressdata.Create(address)

	if err != nil {
		return nil, nil
	}
	// reply := &pb.AddressList{
	// 	Id:      address.ID,
	// 	Address: address.Address,
	// 	Pincode: address.Pincode,
	// 	Type:    address.Type,
	// }
	return &pb.CreateAddressReply{
		Id:      address.ID,
		Address: address.Address,
		Pincode: address.Pincode,
		Type:    address.Type,
	}, nil
}
func (b *Addressbiz) GetAddress(ctx context.Context, req *pb.GetAddressRequest) (*pb.GetAddressReply, error) {
	address := &models.Address{
		ID: req.Id,
	}

	address, err := b.addressdata.GetByID(address.ID)
	if err != nil {
		return nil, nil
	}
	return &pb.GetAddressReply{
		Id:      address.ID,
		Address: address.Address,
		Pincode: address.Pincode,
		Type:    address.Type,
	}, nil
}

func (b *Addressbiz) UpdateAddress(ctx context.Context, req *pb.UpdateAddressRequest) (*pb.UpdateAddressReply, error) {
	address := &models.Address{
		ID:      req.Id,
		Address: req.Address,
		Pincode: req.Pincode,
		Type:    req.Type,
	}
	err := b.addressdata.Update(address)
	if err != nil {
		return nil, nil
	}

	return &pb.UpdateAddressReply{
		Id:      address.ID,
		Address: address.Address,
		Pincode: address.Pincode,
		Type:    address.Type,
	}, nil
}
func (b *Addressbiz) DeleteAddress(ctx context.Context, req *pb.DeleteAddressRequest) (*pb.DeleteAddressReply, error) {
	address := &models.Address{
		ID: req.Id,
	}
	err := b.addressdata.Delete(address.ID)
	if err != nil {
		return nil, nil
	}

	return &pb.DeleteAddressReply{
		Message: "address deleted successfully",
	}, nil
}

// func (b *Addressbiz) ListAddresses(ctx context.Context, req *pb.ListAddressRequest) (*pb.ListAddressReply, error) {
// 	// address := &models.Address{
// 	// 	ID: req.
// 	// }
// 	addresses, err := b.addressdata.List()
// 	if err != nil {
// 		return nil, nil
// 	}
// 	var address []*pb.AddressList
// 	for _, value := range addresses {
// 		address = append(address, &pb.AddressList{
// 			Id:      value.ID,
// 			Address: value.Address,
// 			Pincode: value.Pincode,
// 			Type:    value.Type,
// 		})
// 	}
// 	return &pb.ListAddressReply{
// 		Address: address,
// 	}, nil
// }
