package service

import (
	"log"
	"os"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/JasurbekUz/ToDo-service/genproto"
)

var client pb.TodoServiceClient

func TestMain(m *testing.M) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client = pb.NewTodoServiceClient(conn)

	os.Exit(m.Run())
}
