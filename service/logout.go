package service
import (
	rpc "cloudAuthServer/proto"
	"cloudAuthServer/session"

	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service)LogOut(ctx context.Context,sess *rpc.Session) (*rpc.Message,error) {
	isSess,err := session.Exist(string(sess.Session))
	if isSess {
		return nil,status.Error(codes.InvalidArgument,"not login this session")
	}else if err != nil {
		log.Println(err.Error())
		return nil, status.Error(codes.Internal,err.Error())
	}

	err = session.Pop(string(sess.Session))
	if err != nil {
		log.Println(err.Error())
		return nil, status.Error(codes.Internal,err.Error())
	}
	return &rpc.Message{
		StatusCode: 200,
		Message: "logout",
	},nil
}


