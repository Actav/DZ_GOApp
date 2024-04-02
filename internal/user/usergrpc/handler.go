package usergrpc

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.com/robotomize/gb-golang/homework/03-02-umanager/pkg/pb"
)

var _ pb.UserServiceServer = (*Handler)(nil)

func New(usersRepository usersRepository, timeout time.Duration) *Handler {
	return &Handler{usersRepository: usersRepository, timeout: timeout}
}

type Handler struct {
	pb.UnimplementedUserServiceServer
	usersRepository usersRepository
	timeout         time.Duration
}

func (h Handler) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	// TODO implement me
	panic("implement me")
}

func (h Handler) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.User, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	// TODO implement me
	panic("implement me")
}

func (h Handler) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	// TODO implement me
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}

func (h Handler) DeleteUser(
	ctx context.Context,
	in *pb.DeleteUserRequest,
) (*pb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	// TODO implement me
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}

func (h Handler) ListUsers(
	ctx context.Context,
	in *pb.Empty,
) (*pb.ListUsersResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	// TODO implement me
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}
