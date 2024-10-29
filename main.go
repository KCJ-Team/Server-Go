package main

import (
	"log"
	"net"
	"server-go/src/player"
	"server-go/src/player/pb"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// TLS 설정 적용된 gRPC 서버 생성
	grpcServer := grpc.NewServer()

	// gRPC 서비스 등록
	pb.RegisterPlayerServiceServer(grpcServer, &player.Server{})

	log.Println("Starting secure gRPC server on port 50051 with TLS...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
