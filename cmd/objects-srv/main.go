package main

import (
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	h "try_grpc/internal/handler"
	"try_grpc/internal/memstore"
	"try_grpc/pkg/pb"
)

func main() {
	memStore := memstore.New()
	handler := h.NewHandler(memStore, 10*time.Second)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s) // этот код нужен для дебаггинга
	pb.RegisterLocationServiceServer(s, handler)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
