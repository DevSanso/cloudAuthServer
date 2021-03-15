package service


import (
	rpc "cloudAuthServer/proto"
	"cloudAuthServer/db"
	"cloudAuthServer/session"

	"context"
	"log"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func (s *Service)SetContainerId(ctx context.Context,c *rpc.Container) (*rpc.Message, error) {
	isSess,err := session.Exist(string(c.Session.Session))
	if isSess {
		return nil,status.Error(codes.InvalidArgument,"not login this session")
	}else if err != nil {
		log.Println(err.Error())
		return nil, status.Error(codes.Internal,err.Error())
	}
	rawId,getErr := session.Get(string(c.Session.Session))
	if getErr != nil {
		log.Println(getErr.Error())
		return nil,status.Error(codes.Internal,"Internal session error")
	}
	uid,convErr := strconv.Atoi(rawId)
	if convErr != nil {
		log.Println(convErr.Error())
		return nil,status.Error(codes.Internal,"Internal convert error")
	}

	setErr := db.SetContainerId(ctx,uid,int(c.ContainerId))
	if setErr != nil {
		log.Panicln(setErr.Error())
		return nil,status.Error(codes.Internal,setErr.Error())
	}
	
	return &rpc.Message{
		Message: "done",
		StatusCode: 200,
	},nil
}