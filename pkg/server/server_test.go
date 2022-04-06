package server

import (
	"context"
	"testing"

	assert "github.com/stretchr/testify/require"

	pbv1alpha1 "github.com/fengye87/hello/pkg/pb/v1alpha1"
)

func TestGreet(t *testing.T) {
	s := New()

	tests := []struct {
		req *pbv1alpha1.GreetRequest
		rep *pbv1alpha1.GreetReply
	}{{
		req: &pbv1alpha1.GreetRequest{
			Long: false,
		},
		rep: &pbv1alpha1.GreetReply{
			Message: "Hello!",
		},
	}, {
		req: &pbv1alpha1.GreetRequest{
			Long: true,
		},
		rep: &pbv1alpha1.GreetReply{
			Message: "Hello, World!",
		},
	}}

	for _, tt := range tests {
		rep, err := s.Greet(context.Background(), tt.req)
		assert.NoError(t, err)
		assert.Equal(t, tt.rep.Message, rep.Message)
	}
}
