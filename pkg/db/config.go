package db

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type Config struct {
	Driver             string
	Name               string
	Host               string
	Port               int
	Username           string
	Password           string
	Query              string
	MaxIdleConnections int `mapstructure:"max_idle_connections"`
	MaxOpenConnections int `mapstructure:"max_open_connections"`
	ConnMaxLifetimeMin int `mapstructure:"conn_max_lifetime_min"`
	ConnMaxIdleTime    int `mapstructure:"conn_max_idle_time"`
	Migration          MigrationConfig
}

type MigrationConfig struct {
	Path string
}

func (cfg Config) URL() string {
	var buf bytes.Buffer

	buf.WriteString(cfg.Driver)
	buf.WriteString("://")

	// [username[:password]@]
	if len(cfg.Username) > 0 {
		buf.WriteString(cfg.Username)

		if len(cfg.Password) > 0 {
			buf.WriteByte(':')
			buf.WriteString(cfg.Password)
		}

		buf.WriteByte('@')
	}

	// [host[:port]]
	if len(cfg.Host) > 0 {
		buf.WriteString(cfg.Host)

		if cfg.Port > 0 {
			buf.WriteByte(':')
			buf.WriteString(strconv.Itoa(cfg.Port))
		}
	}

	// /dbname
	buf.WriteByte('/')
	buf.WriteString(cfg.Name)

	// ?query=value
	if len(cfg.Query) > 0 {
		buf.WriteByte('?')
		buf.WriteString(cfg.Query)
	}

	return buf.String()
}

func (cfg Config) ConnectionString() string {
	var kvs []string

	if len(cfg.Username) > 0 {
		kvs = append(kvs, fmt.Sprintf("user=%s", cfg.Username))
	}

	if len(cfg.Password) > 0 {
		kvs = append(kvs, fmt.Sprintf("password=%s", cfg.Password))
	}

	if len(cfg.Host) > 0 {
		kvs = append(kvs, fmt.Sprintf("host=%s", cfg.Host))
	}

	if cfg.Port > 0 {
		kvs = append(kvs, fmt.Sprintf("port=%d", cfg.Port))
	}

	if len(cfg.Name) > 0 {
		kvs = append(kvs, fmt.Sprintf("dbname=%s", cfg.Name))
	}

	if len(cfg.Query) > 0 {
		queries := strings.Split(cfg.Query, "&")
		kvs = append(kvs, queries...)
	}

	return strings.Join(kvs, " ")
}
