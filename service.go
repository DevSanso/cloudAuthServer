package main

import (
	rpc "cloudAuthServer/proto"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	rpc.AuthServiceServer
}

func (s *Service) Login(ctx context.Context, access *rpc.UserAccess) (*rpc.Session, error) {

}
func (s *Service) SignUp(context.Context, *rpc.UserAccess) (*rpc.Session, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (s *Service) Logout(context.Context, *rpc.Session) (*rpc.Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (s *Service) SetContainerId(context.Context, *rpc.Container) (*rpc.Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetContainerId not implemented")
}
func (s *Service) GetContainerId(context.Context, *rpc.Session) (*rpc.Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContainerId not implemented")
}
func (s *Service) DeleteAccount(context.Context, *rpc.Session) (*rpc.Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (s *Service) EmailVail(context.Context, *rpc.Session) (*rpc.Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmailVail not implemented")
}
func (s *Service) IsVailEmail(context.Context, *rpc.Session) (*rpc.Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsVailEmail not implemented")
}
