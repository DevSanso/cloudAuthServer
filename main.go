package main

import (
	"net"
	"log"
	"strconv"
	"google.golang.org/grpc"
)

func main() {
	cfg,cfgErr := newServerArgs()
	if cfgErr != nil {log.Panicln(cfgErr)}

	l,lErr := net.Listen("tcp",cfg.Host + ":" + strconv.Itoa(cfg.Port))
	if lErr != nil {log.Panicln(lErr)}

	grpcServer := grpc.NewServer()
	log.Println("start cloudAuthServer -> ",cfg.Host+":",cfg.Port)
	
	if grpcErr := grpcServer.Serve(l);grpcErr != nil {
		log.Panicln(grpcErr)
	}
}