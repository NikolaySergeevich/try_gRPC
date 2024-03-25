package handler

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	// "gitlab.com/robotomize/gb-golang/homework/03-02-grpc-solve/internal/geo"
	"github.com/NikolaySergeevich/try_gRPC/internal/geo"
	"trygrpc/internal/geo"
	"trygrpc/internal/memstore"
	"trygrpc/pkg/pb"
)

var _ pb.LocationServiceServer = (*Handler)(nil)

func NewHandler(store store, timeout time.Duration) *Handler {
	return &Handler{store: store, timeout: timeout}
}

type Handler struct {
	pb.UnimplementedLocationServiceServer
	store   store
	timeout time.Duration
}

func (h Handler) AddObject(ctx context.Context, object *pb.Object) (*pb.ObjectID, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	if object == nil {
		return nil, status.Error(codes.InvalidArgument, "id is empty")
	}

	h.store.Add(
		memstore.Item{
			ID:   object.Id,
			Name: object.Name,
			Lat:  object.Lat,
			Lon:  object.Lon,
		},
	)
	return &pb.ObjectID{Id: object.Id}, nil
}

func (h Handler) GetObject(ctx context.Context, id *pb.ObjectID) (*pb.Object, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	if id == nil {
		return nil, status.Error(codes.InvalidArgument, "id is empty")
	}

	item, ok := h.store.FindByObjectID(id.GetId())
	if !ok {
		return nil, status.Error(codes.NotFound, "item not found")
	}

	return &pb.Object{
		Id:   item.ID,
		Name: item.Name,
		Lat:  item.Lat,
		Lon:  item.Lon,
	}, nil
}

func (h Handler) DeleteObject(ctx context.Context, id *pb.ObjectID) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	if id == nil {
		return nil, status.Error(codes.InvalidArgument, "id is empty")
	}

	h.store.DeleteByObjectID(id.GetId())
	return &emptypb.Empty{}, nil
}

func (h Handler) ListObjects(ctx context.Context, _ *emptypb.Empty) (*pb.ObjectList, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	items := h.store.FindAll()
	respObjects := make([]*pb.Object, len(items))

	for _, item := range items {
		respObjects = append(
			respObjects, &pb.Object{
				Id:   item.ID,
				Lat:  item.Lat,
				Lon:  item.Lon,
				Name: item.Name,
			},
		)
	}
	return &pb.ObjectList{Objects: respObjects}, nil
}

func (h Handler) CalculateDistance(ctx context.Context, request *pb.DistanceRequest) (*pb.DistanceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "request invalid")
	}

	item, ok := h.store.FindByObjectID(request.GetObjectId())
	if !ok {
		return nil, status.Error(codes.NotFound, "item not found")
	}

	dist := geo.ComputeDistance(item.Lat, item.Lon, request.Lat, request.Lon)

	return &pb.DistanceResponse{Distance: dist}, nil
}
