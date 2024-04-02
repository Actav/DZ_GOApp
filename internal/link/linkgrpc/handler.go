package linkgrpc

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.com/robotomize/gb-golang/homework/03-02-umanager/pkg/pb"
)

var _ pb.LinkServiceServer = (*Handler)(nil)

func New(linksRepository linksRepository, timeout time.Duration) *Handler {
	return &Handler{linksRepository: linksRepository, timeout: timeout}
}

type Handler struct {
	pb.UnimplementedLinkServiceServer
	linksRepository linksRepository
	timeout         time.Duration
}

func (h Handler) GetLinkByUserID(ctx context.Context, id *pb.GetLinksByUserId) (*pb.ListLinkResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (h Handler) mustEmbedUnimplementedLinkServiceServer() {
	// TODO implement me
	panic("implement me")
}

func (h Handler) CreateLink(ctx context.Context, request *pb.CreateLinkRequest) (*pb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	// TODO implement me
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}

func (h Handler) GetLink(ctx context.Context, request *pb.GetLinkRequest) (*pb.Link, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	// TODO implement me
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}

func (h Handler) UpdateLink(ctx context.Context, request *pb.UpdateLinkRequest) (*pb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	// TODO implement me
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}

func (h Handler) DeleteLink(ctx context.Context, request *pb.DeleteLinkRequest) (*pb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	// TODO implement me
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}

func (h Handler) ListLinks(ctx context.Context, request *pb.Empty) (*pb.ListLinkResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	// TODO implement me
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}
