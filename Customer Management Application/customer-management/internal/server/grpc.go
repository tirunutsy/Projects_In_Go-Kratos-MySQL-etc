package server

import (
	v1 "customer-management/api/helloworld/v1"
	v3 "customer-management/api/helloworld/v1/address"
	v2 "customer-management/api/helloworld/v1/phone"
	"customer-management/internal/conf"
	"customer-management/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, customer *service.CustomerService,
	email *service.EmailService, phone *service.PhoneService, address *service.AddressService,
	logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterCustomerServer(srv, customer)
	v1.RegisterEmailServer(srv, email)
	v2.RegisterPhoneServiceServer(srv, phone)
	v3.RegisterAddressServer(srv, address)
	return srv
}
