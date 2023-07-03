package tests

import (
	"RssReader/bussiness/data/schema"
	"RssReader/bussiness/sys/config"
	"RssReader/foundation/docker"
	"bufio"
	"bytes"
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
	"time"
)

// Success and failure markers.
const (
	Success = "\u2713"
	Failed  = "\u2717"
)

// StartDB starts a database instance.
func StartDB() (*docker.Container, error) {
	image := "postgres:13-alpine"
	port := "5432"
	args := []string{"-e", "POSTGRES_PASSWORD=postgres"}

	return docker.StartContainer(image, port, args...)
}

// StopDB stops a running database instance.
func StopDB(c *docker.Container) {
	docker.StopContainer(c.ID)
}

func CreateDB(c *docker.Container) error {
	commands := []string{"psql", "-U", "postgres", "-d", "postgres", "-c", "CREATE DATABASE testing;"}
	if err := docker.ExecuteCommand(c.ID, commands...); err != nil {
		return err
	}
	return nil
}

func DropDB(c *docker.Container) error {
	commands := []string{"psql", "-U", "postgres", "-d", "postgres", "-c", "DROP DATABASE testing;"}
	if err := docker.ExecuteCommand(c.ID, commands...); err != nil {
		return err
	}
	return nil
}

func NewUnit(t *testing.T, db *docker.Container) (*zap.SugaredLogger, *gorm.DB, func()) {
	if err := CreateDB(db); err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		db.Host,
		"postgres",
		"postgres",
		"testing",
		db.Port,
	)
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		t.Fatalf("Opening database connection: %v", err)
	}
	t.Log("waiting for database to be ready ...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := schema.Migrate(config.Database{
		Host:     db.Host,
		User:     "postgres",
		Password: "postgres",
		Dbname:   "testing",
		Port:     db.Port,
	}); err != nil {

		docker.DumpContainerLogs(t, db.ID)
		DropDB(db)
		t.Fatalf("Migrating error: %s", err)
	}

	if err := schema.Seed(ctx, Db); err != nil {
		docker.DumpContainerLogs(t, db.ID)
		DropDB(db)
		t.Fatalf("Sending error: %s", err)
	}

	var buf bytes.Buffer
	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	writer := bufio.NewWriter(&buf)
	log := zap.New(
		zapcore.NewCore(encoder, zapcore.AddSync(writer), zapcore.DebugLevel)).
		Sugar()

	teardown := func() {
		t.Helper()
		sqlDB, err := Db.DB()
		if err != nil {
			log.Fatal(err)
		}
		sqlDB.Close()
		DropDB(db)
		log.Sync()
		writer.Flush()
		fmt.Println("******************** LOGS ********************")
		fmt.Print(buf.String())
		fmt.Println("******************** LOGS ********************")

	}
	return log, Db, teardown
}
