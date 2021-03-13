package service
import (
	rpc "cloudAuthServer/proto"
	"cloudAuthServer/db/document"
	"cloudAuthServer/db"

	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service)LogOut(ctx context.Context,session *rpc.Session) (*rpc.Message,error) {
	
}


