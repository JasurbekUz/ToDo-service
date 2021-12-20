package service

import (
	"context"
	"testing"

	pb "github.com/JasurbekUz/ToDo-service/genproto"
)

var (
	id   string
	todo *pb.Todo
	err  error
)

func TestTodoService_Create(t *testing.T) {
	todo, err = client.Create(context.Background(), &pb.Todo{
		Assignee: "assignee_3",
		Title:    "title_3",
		Summary:  "summary_3",
		Deadline: "2021-12-20",
		Status:   "status_3",
	})

	if err != nil {
		t.Error(err)
	}

	id = todo.Id
}

func TestTodoService_Get(t *testing.T) {
	_, err := client.Get(context.Background(), &pb.ByIdReq{
		Id: id,
	})
	if err != nil {
		t.Error(err)
	}
}

func TestTodoService_Update(t *testing.T) {
	_, err := client.Update(context.Background(), &pb.Todo{
		Id:       id,
		Assignee: "assignee_edited",
		Title:    "title_edited",
		Summary:  "summary_edited",
		Deadline: "2021-12-20",
		Status:   "status_edited",
	})
	if err != nil {
		t.Error(err)
	}
}

func TestTodoService_List(t *testing.T) {
	_, err := client.List(context.Background(), &pb.ListReq{
		Page:  2,
		Limit: 1,
	})
	if err != nil {
		t.Error(err)
	}
}

func TestTodoService_ListOverdue(t *testing.T) {
	_, err := client.ListOverdue(context.Background(), &pb.ListTime{
		ListPage: &pb.ListReq{
			Page:  2,
			Limit: 1,
		},
		ToTime: "2021-12-19",
	})
	if err != nil {
		t.Error(err)
	}
}

func TestTodoService_Delete(t *testing.T) {
	_, err := client.Delete(context.Background(), &pb.ByIdReq{
		Id: id,
	})
	if err != nil {
		t.Error(err)
	}
}
