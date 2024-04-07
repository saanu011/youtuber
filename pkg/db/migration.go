package db

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres" // postgres
	_ "github.com/golang-migrate/migrate/v4/source/file"       // file migration
	_ "github.com/lib/pq"                                      // psql lib
)

func RunMigrations(conf Config) error {
	fmt.Printf("Ensuring database migration")

	m, err := newMigrate(conf)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			return nil
		}

		return err
	}

	return nil
}

func RollbackMigration(conf Config) error {
	m, err := newMigrate(conf)
	if err != nil {
		return err
	}

	err = m.Steps(-1)
	if err != nil {
		if err == migrate.ErrNoChange {
			return nil
		}

		return err
	}

	return nil
}

func CreateMigration(filename string, conf Config) error {
	if len(filename) == 0 {
		return errors.New("filename is not provided")
	}

	u, err := url.Parse(conf.Migration.Path)
	if err != nil {
		return err
	}

	path := u.Path

	timestamp := time.Now().Unix()
	upMigrationFilePath := fmt.Sprintf("%s/%d_%s.up.sql", path, timestamp, filename)
	downMigrationFilePath := fmt.Sprintf("%s/%d_%s.down.sql", path, timestamp, filename)

	if err := createFile(upMigrationFilePath); err != nil {
		return err
	}

	fmt.Printf("created %s\n", upMigrationFilePath)

	if err := createFile(downMigrationFilePath); err != nil {
		os.Remove(upMigrationFilePath)
		return err
	}

	fmt.Printf("created %s\n", downMigrationFilePath)

	return nil
}

func newMigrate(conf Config) (*migrate.Migrate, error) {
	return migrate.New(conf.Migration.Path, conf.URL())
}

func createFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	err = f.Close()

	return err
}
