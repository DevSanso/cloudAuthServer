package service



import (
	rpc "cloudAuthServer/proto"
	"cloudAuthServer/db/document"
	"cloudAuthServer/db"
	"cloudAuthServer/session"
	"cloudAuthServer/utils"

	"context"

	"time"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

)

var SessionAllocTimeout = time.Hour * 1


func (s *Service)Login(ctx context.Context,user *rpc.UserAccess) (*rpc.SessionResult,error) {
	tempSess :=string(utils.Make256HashSalt(user.Email,user.Passwd))
	isExt,err :=  session.Exist(tempSess)
	if isExt {
		return nil,status.Error(codes.InvalidArgument,"already login user")
	}else if err != nil {
		return nil,status.Error(codes.Internal,err.Error())
	}

	
	isExt,err = db.ExistUser(ctx,&document.User{
		user.GetEmail(),
		user.GetPasswd(),
		false,
	})

	if !isExt {
		return nil,status.Error(codes.InvalidArgument,"not register with user")
	}else if err != nil {
		return nil,status.Error(codes.Internal,err.Error())
	}

	expire := time.Now().Add(SessionAllocTimeout)
	err = session.SetTimeOut(tempSess,"",SessionAllocTimeout)
	if err != nil {
		return nil,status.Errorf(codes.Internal,err.Error())
	}

	var res = &rpc.SessionResult{
		Message : &rpc.Message {
			StatusCode: 200,
			Message : fmt.Sprint("Expire:",expire.Format(time.RFC850)),
		},
		Session : &rpc.Session{
			Session : []byte(tempSess),
		},
	}
	return res,nil

}