package handler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"trygrpc/pkg/pb"
)

var _ pb.LocationServiceServer = (*Handler)(nil)

func NewHandler() *Handler {
	return &Handler{}
}

type Handler struct {
	pb.UnimplementedLocationServiceServer
}

func (h Handler) AddObject(ctx context.Context, object *pb.Object) (*pb.ObjectID, error) {
	// TODO implement me
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

func (h Handler) GetObject(ctx context.Context, id *pb.ObjectID) (*pb.Object, error) {
	// TODO implement me
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

func (h Handler) DeleteObject(ctx context.Context, id *pb.ObjectID) (*emptypb.Empty, error) {
	// TODO implement me
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

func (h Handler) ListObjects(ctx context.Context, empty *emptypb.Empty) (*pb.ObjectList, error) {
	// TODO implement me
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

func (h Handler) CalculateDistance(ctx context.Context, request *pb.DistanceRequest) (*pb.DistanceResponse, error) {
	// TODO implement me
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}
