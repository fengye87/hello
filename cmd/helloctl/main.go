package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pbv1alpha1 "github.com/fengye87/hello/pkg/pb/v1alpha1"
)

func main() {
	var server = "localhost:8080"
	flag.StringVar(&server, "server", server, "Server address")
	var long bool
	flag.BoolVar(&long, "long", long, "Long")
	flag.Parse()

	conn, err := grpc.Dial(server, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pbv1alpha1.NewHelloClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := pbv1alpha1.GreetRequest{Long: long}
	rep, err := client.Greet(ctx, &req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", rep.GetMessage())
}
