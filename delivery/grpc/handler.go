package grpc

import (
	"context"

	"google.golang.org/grpc"

	"bitbucket.org/yesboss/sharingan/controller"
	pb "bitbucket.org/yesboss/sharingan/proto"
)

type Handler struct {
	UserController controller.UserController
}

func (h *Handler) GetResourceId(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	// a := h.RoleController.Get("a")
	return &pb.Response{ResourceId: "dummy"}, nil
}

func New(c controller.UserController) *grpc.Server {
	s := grpc.NewServer()
	handler := Handler{UserController: c}

	pb.RegisterDataServer(s, &handler)

	return s
}
