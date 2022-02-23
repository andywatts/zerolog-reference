package main

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"zerolog-reference/pkg/pb"
)

type blahServer struct {
	pb.UnimplementedBlahServiceServer
}

func newServer() *blahServer {
	s := &blahServer{}
	return s
}

func (s *blahServer) GetBlah(ctx context.Context, request *pb.MyRequest) (*pb.MyResponse, error) {
	return &pb.MyResponse{Response: request.Request + "!"}, nil
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Print("hello world")

	lis, _ := net.Listen("tcp", "localhost:8080")

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterBlahServiceServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
