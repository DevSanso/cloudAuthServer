package service

import (
	
	rpc "cloudAuthServer/proto"
	"cloudAuthServer/db"
	"cloudAuthServer/session"

	"context"
	"log"
	"strconv"

	"github.com/tidwall/buntdb"

	"google.golang.org/grpc/codes"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/status"
)

func (s *Service)GetContainer(ctx context.Context,sess *rpc.Session) (*rpc.Message,error) {
	userIdStr,err := session.Get(string(sess.Session))
	if err == buntdb.ErrNotFound {
		return nil,status.Error(codes.InvalidArgument,err.Error())
	}else if err != nil {
		log.Println(err.Error())
		return nil,status.Error(codes.Internal,err.Error())
	}
	userId, convErr := strconv.Atoi(userIdStr)
	if convErr != nil {
		panic(convErr.Error())
	}
	cId,gErr := db.GetContainer(ctx,userId)
	if gErr == mongo.ErrNoDocuments {
		return nil,status.Error(codes.InvalidArgument,err.Error())
	}else if gErr != nil {
		return nil,status.Error(codes.Internal,gErr.Error())
	}
	return &rpc.Message{
		StatusCode: 200,
		Message: strconv.Itoa(cId),
	},nil
}