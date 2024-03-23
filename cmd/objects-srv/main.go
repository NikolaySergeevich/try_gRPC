package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	h "try_grpc/internal/handler"
	"trygrpc/pkg/pb"
)

func main() {
	handler := h.NewHandler()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s) // этот код нужен для дебаггинга
	pb.RegisterLocationServiceServer(s, handler)

	slog.Info(fmt.Sprintf("grpc serve %s", ":50051"))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
