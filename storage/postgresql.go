package storage

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mtfelian/test_task_3/config"
	"github.com/mtfelian/test_task_3/models"
	"github.com/spf13/viper"
)

// DriverNamePostgres holds driver name for PostgreSQL
const DriverNamePostgres = "postgres"

// PostgresDSNTests returns a postgresql DSN (DB access config string) for testing purposes
func PostgresDSNTests() string {
	return `host=localhost port=5432 user=postgres dbname=test_task_3 sslmode=disable client_encoding=utf8`
}

// DB provides to DB access
func (s Postgres) DB() *gorm.DB {
	if s.db != nil {
		return s.db
	}

	fmt.Println("connecting to DB:", viper.GetString(config.DSN))
	newDB, err := gorm.Open(viper.GetString(config.DBDriver), viper.GetString(config.DSN))
	if err != nil {
		panic(err)
	}
	s.db = newDB

	//s.db.LogMode(true)
	s.db.SingularTable(true)
	return s.db
}

// NewPostgres creates a new PostgreSQL-based storage via GORM and default/loaded DB settings
func NewPostgres() Postgres {
	p := Postgres{}
	p.db = p.DB()
	return p
}

// Postgres implements Keeper for postgres
type Postgres struct{ db *gorm.DB }

// Get a model by ID
func (s Postgres) Get(_type, data string) (*models.Param, error) {
	var param models.Param
	err := s.DB().First(&param, `"type" = ? AND "data" = ?`, _type, data).Error
	if err == gorm.ErrRecordNotFound {
		return &models.Param{Type: _type, Data: data}, nil
	}
	if err != nil {
		return nil, err
	}
	return &param, nil
}

// Close the link to storage
func (s Postgres) Close() error { return s.db.Close() }
