package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURL(t *testing.T) {
	conf := Config{
		Driver:             "postgres",
		Host:               "localhost",
		Port:               5432,
		Name:               "youtuber_test",
		Username:           "admin",
		Password:           "123456",
		Query:              "sslmode=disable&timezone=UTC",
		MaxIdleConnections: 2,
		MaxOpenConnections: 50,
		ConnMaxLifetimeMin: 10,
		ConnMaxIdleTime:    3000,
		Migration: MigrationConfig{
			Path: "file://db/migrations",
		},
	}

	expectedURL := "postgres://admin:123456@localhost:5432/youtuber_test?sslmode=disable&timezone=UTC"
	actualURL := conf.URL()
	assert.Equal(t, expectedURL, actualURL)
}

func TestConnectionString(t *testing.T) {
	conf := Config{
		Driver:             "postgres",
		Host:               "localhost",
		Port:               5432,
		Name:               "youtuber_test",
		Username:           "admin",
		Password:           "123456",
		Query:              "sslmode=disable&timezone=UTC",
		MaxIdleConnections: 2,
		MaxOpenConnections: 50,
		ConnMaxLifetimeMin: 10,
		ConnMaxIdleTime:    3000,
		Migration: MigrationConfig{
			Path: "file://db/migrations",
		},
	}

	expectedConnectionString := "user=admin password=123456 host=localhost port=5432 dbname=youtuber_test sslmode=disable timezone=UTC"
	actualConnectionString := conf.ConnectionString()
	assert.Equal(t, expectedConnectionString, actualConnectionString)
}
