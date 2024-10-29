package player

import (
	"context"
	"log"

	pb "server-go/src/player/pb"
)

// Server 구조체는 gRPC 인터페이스를 구현합니다.
type Server struct {
	pb.UnimplementedPlayerServiceServer
}

// SendPlayerPosition 메서드는 플레이어의 위치를 수신합니다.
func (s *Server) SendPlayerPosition(ctx context.Context, req *pb.RequestPlayerPosition) (*pb.Response, error) {
	log.Printf("Received position from PlayerID: %s, X: %f, Y: %f, Z: %f", req.PlayerId, req.X, req.Y, req.Z)
	return &pb.Response{Message: "Position received successfully!"}, nil
}
