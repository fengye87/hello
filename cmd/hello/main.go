package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"

	pbv1alpha1 "github.com/fengye87/hello/pkg/pb/v1alpha1"
	"github.com/fengye87/hello/pkg/server"
)

func main() {
	var bind = "0.0.0.0:8080"
	flag.StringVar(&bind, "bind", bind, "Bind address")
	flag.Parse()

	l, err := net.Listen("tcp", bind)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	m := cmux.New(l)
	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpL := m.Match(cmux.Any())

	grpcS := grpc.NewServer()
	grpcM := runtime.NewServeMux()

	helloS := server.New()
	pbv1alpha1.RegisterHelloServer(grpcS, helloS)
	if err := pbv1alpha1.RegisterHelloHandlerServer(context.Background(), grpcM, helloS); err != nil {
		log.Fatalf("failed to register hello handler server: %v", err)
	}

	go grpcS.Serve(grpcL)
	go http.Serve(httpL, grpcM)

	log.Printf("server listening at %v", l.Addr())
	if err := m.Serve(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
