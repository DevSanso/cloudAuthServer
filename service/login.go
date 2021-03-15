package service



import (
	rpc "cloudAuthServer/proto"
	"cloudAuthServer/db/document"
	"cloudAuthServer/db"
	"cloudAuthServer/session"
	"cloudAuthServer/utils"

	"context"
	"log"
	"time"
	"strconv"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
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
		log.Println(err.Error())
		return nil,status.Error(codes.Internal,err.Error())
	}

	
	uId,getErr := db.GetUserId(ctx,&document.User{
		user.GetEmail(),
		user.GetPasswd(),
		false,
	})

	if getErr == mongo.ErrNoDocuments {
		return nil,status.Error(codes.InvalidArgument,"not register with user")
	}else if err != nil {
		log.Println(err.Error())
		return nil,status.Error(codes.Internal,err.Error())
	}

	expire := time.Now().Add(SessionAllocTimeout)
	err = session.SetTimeOut(tempSess,strconv.Itoa(uId),SessionAllocTimeout)
	if err != nil {
		log.Println(err.Error())
		return nil,status.Errorf(codes.Internal,err.Error())
	}

	var res = &rpc.SessionResult{
		Message : &rpc.Message {
			StatusCode: 200,
			Message : fmt.Sprint("timeout:",expire.Format(time.RFC850)),
		},
		Session : &rpc.Session{
			Session : []byte(tempSess),
		},
	}
	return res,nil

}