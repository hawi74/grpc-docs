package main

import (
	context "context"
	"log"
	"net"

	"github.com/hawi74/grpc-docs/generated"
	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type server struct{
	generated.UnimplementedTodoServiceServer
}

// CreateTodo implements generated.TodoServiceServer
func (*server) CreateTodo(context.Context, *generated.Todo) (*emptypb.Empty, error) {
	log.Print("Serving /CreateTodo")
	return &emptypb.Empty{}, nil
}

// SearchTodo implements generated.TodoServiceServer
func (*server) SearchTodo(context.Context, *generated.SearchTodoRequest) (*generated.SearchTodoResult, error) {
	log.Print("Serving /SearchTodo")
	return &generated.SearchTodoResult{
		Todos: []*generated.Todo{
			{
				Id: "1",
				Title: "foobar",
				Done: false,
			},
			{
				Id: "2",
				Title: "baz",
				Done:  true,
			},
			{
				Id: "3",
				Title: "????",
				Done: false,
			},
		},
		Total: 3,
		TotalPages: 1,
	}, nil
}

var _ generated.TodoServiceServer = (*server)(nil)

func main() {

	grpcServer := grpc.NewServer()

	generated.RegisterTodoServiceServer(grpcServer, &server{})

	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("cannot listen: %v", err)
	}

	log.Print("listening on :8000")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("cannot serve: %v", err)
	}
}
