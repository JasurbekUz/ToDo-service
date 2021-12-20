package postgres

import (
	"testing"

	"github.com/JasurbekUz/ToDo-service/config"
	pb "github.com/JasurbekUz/ToDo-service/genproto"
	"github.com/JasurbekUz/ToDo-service/pkg/db"
	"github.com/JasurbekUz/ToDo-service/storage/repo"

	"github.com/stretchr/testify/suite"
)

type TodoRepositoryTestSuite struct {
	suite.Suite
	CleanupFunc func()
	Repository  repo.TodoStorageI
}

func (suite *TodoRepositoryTestSuite) SetupSuite() {
	pgPool, cleanup := db.ConnectDBForSuite(config.Load())

	suite.Repository = NewTodoRepo(pgPool)
	suite.CleanupFunc = cleanup
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *TodoRepositoryTestSuite) TestTodoCRUD() {
	id := "0d512776-60ed-4980-b8a3-6904a2234fd9"

	todo := pb.Todo{
		Id:       id,
		Assignee: "assignee",
		Title:    "title_1",
		Summary:  "summary_1",
		Deadline: "2021-12-15T14:12:14Z",
		Status:   "active",
	}

	_ = suite.Repository.Delete(id)

	// Creating_Test_Part
	todo, err := suite.Repository.Create(todo)
	suite.Nil(err)

	// Getting_Test_Part
	getTodo, err := suite.Repository.Get(todo.Id)
	suite.Nil(err)
	suite.NotNil(getTodo, "todo must not be nil")
	suite.Equal(todo.Assignee, getTodo.Assignee, "assigne must match")

	// Updating_Test_Part
	todo.Assignee = "assignee_1"
	updatedTodo, err := suite.Repository.Update(todo)
	suite.Nil(err)

	// Updated_Getting_Test_Part
	getTodo, err = suite.Repository.Get(id)
	suite.Nil(err)
	suite.NotNil(getTodo)
	suite.Equal(todo.Assignee, updatedTodo.Assignee)

	// List_Test_Part
	listTodos, _, err := suite.Repository.List(1, 2)
	suite.Nil(err)
	suite.NotEmpty(listTodos)
	suite.Equal(todo.Id, listTodos[0].Id)

	// Deleting_Test_Part
	err = suite.Repository.Delete(id)
	suite.Nil(err)
}

func (suite *TodoRepositoryTestSuite) TearDownSuite() {
	suite.CleanupFunc()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestUserTodoTestSuite(t *testing.T) {
	suite.Run(t, new(TodoRepositoryTestSuite))
}
