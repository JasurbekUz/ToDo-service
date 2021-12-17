package postgres

import (
	"reflect"
	"testing"
	"time"

	pb "github.com/JasurbekUz/ToDo-service/genproto"
)

var (
	Id_1 = "f6b409cc-8520-4532-4d80-a6c7a6092a8c"
	Id_2 = "5cd22b4d-29fd-4a78-4cc6-cd3af539c363"
)

func TestTodoRepo_Create(t *testing.T) {
	tests := []struct {
		name    string
		input   pb.Todo
		want    pb.Todo
		wantErr bool
	}{
		{
			name: "successful",
			input: pb.Todo{
				Id:       Id_1,
				Assignee: "assignee_1",
				Title:    "title_1",
				Summary:  "summary_1",
				Deadline: "2021-12-15T14:12:14Z",
				Status:   "active",
			},
			want: pb.Todo{
				Id:       Id_1,
				Assignee: "assignee_1",
				Title:    "title_1",
				Summary:  "summary_1",
				Deadline: "2021-12-15T14:12:14Z",
				Status:   "active",
			},
			wantErr: false,
		},
		{
			name: "successful",
			input: pb.Todo{
				Id:       Id_2,
				Assignee: "assignee_2",
				Title:    "title_2",
				Summary:  "summary_2",
				Deadline: "2021-12-18T18:00:10Z",
				Status:   "active",
			},
			want: pb.Todo{
				Id:       Id_2,
				Assignee: "assignee_2",
				Title:    "title_2",
				Summary:  "summary_2",
				Deadline: "2021-12-18T18:00:10Z",
				Status:   "active",
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := pgRepo.Create(tc.input)
			if err != nil {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.wantErr, err)
			}

			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTodoRepo_Get(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    pb.Todo
		wantErr bool
	}{
		{
			name:  "successful",
			input: Id_1,
			want: pb.Todo{
				Id:       Id_1,
				Assignee: "assignee_1",
				Title:    "title_1",
				Summary:  "summary_1",
				Deadline: "2021-12-15T14:12:14Z",
				Status:   "active",
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := pgRepo.Get(tc.input)
			if err != nil {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.wantErr, err)
			}
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTodoRepo_List(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			page, limit int64
		}
		want    []*pb.Todo
		wantErr bool
	}{
		{
			name: "succesful",
			input: struct {
				page, limit int64
			}{
				page:  1,
				limit: 2,
			},
			want: []*pb.Todo{
				{
					Id:       Id_1,
					Assignee: "assignee_1",
					Title:    "title_1",
					Summary:  "summary_1",
					Deadline: "2021-12-15T14:12:14Z",
					Status:   "active",
				},
				{
					Id:       Id_2,
					Assignee: "assignee_2",
					Title:    "title_2",
					Summary:  "summary_2",
					Deadline: "2021-12-18T18:00:10Z",
					Status:   "active",
				},
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, count, err := pgRepo.List(tc.input.page, tc.input.limit)
			if err != nil {
				t.Fatalf("%s: expected: %v, got: %v, count: %d", tc.name, tc.wantErr, err, count)
			}
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v, count: %d", tc.name, tc.want, got, count)
			}
		})
	}
}

func TestTodoRepo_Update(t *testing.T) {
	tests := []struct {
		name    string
		input   pb.Todo
		want    pb.Todo
		wantErr bool
	}{
		{
			name: "successful",
			input: pb.Todo{
				Id:       Id_2,
				Assignee: "assignee_2",
				Title:    "title_2",
				Summary:  "summary_2",
				Deadline: "2021-12-15T14:12:14Z",
				Status:   "active",
			},
			want: pb.Todo{
				Id:       Id_2,
				Assignee: "assignee_2",
				Title:    "title_2",
				Summary:  "summary_2",
				Deadline: "2021-12-15T14:12:14Z",
				Status:   "active",
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := pgRepo.Update(tc.input)
			if err != nil {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.wantErr, err)
			}

			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTodoRepo_Delete(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    error
		wantErr bool
	}{
		{
			name:    "successful",
			input:   Id_2,
			want:    nil,
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := pgRepo.Delete(tc.input)
			if err != nil {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.wantErr, err)
			}
		})
	}
}

func TestTodoRepo_ListOverdue(t *testing.T) {
	fmtTime := "2006-01-02"
	toTime, err := time.Parse(fmtTime, "2021-12-10")
	if err != nil {
		t.Fatal("failed to time parse", err)
	}
	tests := []struct {
		name  string
		input struct {
			time        time.Time
			page, limit int64
		}
		want    []*pb.Todo
		wantErr bool
	}{
		{
			name: "succesful",
			input: struct {
				time        time.Time
				page, limit int64
			}{
				time:  toTime,
				page:  1,
				limit: 2,
			},
			want: []*pb.Todo{
				{
					Id:       Id_1,
					Assignee: "assignee_1",
					Title:    "title_1",
					Summary:  "summary_1",
					Deadline: "2021-12-15T14:12:14Z",
					Status:   "active",
				},
				{
					Id:       Id_2,
					Assignee: "assignee_2",
					Title:    "title_2",
					Summary:  "summary_2",
					Deadline: "2021-12-18T18:00:10Z",
					Status:   "active",
				},
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, count, err := pgRepo.ListOverdue(tc.input.time, tc.input.page, tc.input.limit)
			if err != nil {
				t.Fatalf("%s: expected: %v, got: %v, count: %d", tc.name, tc.wantErr, err, count)
			}
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v, count: %d", tc.name, tc.want, got, count)
			}
		})
	}
}
