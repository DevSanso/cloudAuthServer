package service


import (
	rpc "cloudAuthServer/proto"
	"cloudAuthServer/db/document"
	"cloudAuthServer/db"

	"context"
	"log"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var AlreadyExistUserErr = errors.New("AlreadyExistUserError")

func (s *Service) SignUp(ctx context.Context,userInfo *rpc.UserAccess) (*rpc.SessionResult, error) {
	
	user := &document.User {
		userInfo.GetEmail(),
		userInfo.GetPasswd(),
		false,
	}

	isExist,err := db.ExistUser(ctx,user)

	if err != nil {
		log.Println(err.Error())
		return nil,status.Errorf(codes.Internal,err.Error())
	}
	if isExist {return nil,status.Errorf(codes.AlreadyExists,AlreadyExistUserErr.Error())}

	err = db.RegisterUser(ctx,user)
	if err != nil {return nil,status.Errorf(codes.InvalidArgument,err.Error())}
	ret := &rpc.SessionResult{}
	ret.Message.Message = "done"
	ret.Message.StatusCode = 200
	//ret.Session.Session = ""
	return ret, nil
}