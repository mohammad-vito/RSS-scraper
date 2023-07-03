package schema

import (
	"RssReader/bussiness/sys/config"
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
	"os"
)

func Migrate(conf config.Database) error {
	m, err := migrate.New(
		"file://bussiness/data/schema/migrations",
		fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable",
			conf.User,
			conf.Password,
			conf.Host,
			conf.Port,
			conf.Dbname,
		))
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = m.Up()
	if err != nil {
		fmt.Printf("migration up: %v", err)
		return err
	}
	return nil
}

func Seed(ctx context.Context, db *gorm.DB) error {
	sqlBytes, err := os.ReadFile("bussiness/data/schema/sql/seed.sql")
	if err != nil {
		return err
	}
	db.Exec(string(sqlBytes))
	return nil
}
