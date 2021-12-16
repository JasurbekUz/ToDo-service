package postgres

import (
	"log"
	"os"
	"testing"

	"github.com/JasurbekUz/ToDo-service/config"
	"github.com/JasurbekUz/ToDo-service/pkg/db"
	"github.com/JasurbekUz/ToDo-service/pkg/logger"
)

var pgRepo *todoRepo

func TestMain(m *testing.M) {
	cfg := config.Load()

	connDB, err := db.ConnectionToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	pgRepo = NewTodoRepo(connDB)

	os.Exit(m.Run())
}
