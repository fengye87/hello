package server

import (
	"context"

	pbv1alpha1 "github.com/fengye87/hello/pkg/pb/v1alpha1"
)

func New() pbv1alpha1.HelloServer {
	return &server{}
}

type server struct {
	pbv1alpha1.UnimplementedHelloServer
}

func (s *server) Greet(ctx context.Context, req *pbv1alpha1.GreetRequest) (*pbv1alpha1.GreetReply, error) {
	var message = "Hello!"
	if req.Long {
		message = "Hello, World!"
	}

	rep := pbv1alpha1.GreetReply{
		Message: message,
	}
	return &rep, nil
}
