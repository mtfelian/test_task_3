package storage

import (
	"fmt"
	"strings"
	"sync"
	"time"

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
	return `host=localhost port=5432 user=postgres dbname=test_db sslmode=disable client_encoding=utf8`
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

// GenerateTestEntries with given amount and INSERTion bulkSize just for testing purposes
func (s Postgres) GenerateTestEntries(amount, bulkSize uint) error {
	t1 := time.Now()
	fmt.Println("Generating test data...")
	defer func() { fmt.Printf("Generation OK. Elapsed time: %.3fs\n", time.Now().Sub(t1).Seconds()) }()

	// such generation has a very small chance of some combination of ID and Addr will not be added
	ts := time.Now()
	for i := uint(0); i < amount; i += bulkSize {
		if i%10000 == 0 {
			fmt.Printf("Completed %.2f%%, elapsed %.2fs\r", float64(i)/float64(amount)*100,
				time.Now().Sub(ts).Seconds())
		}

		var params []models.Param
		for j := uint(0); j < bulkSize; j++ {
			// generate model fields
			params = append(params, models.Param{})
		}
		if err := s.AddMultiple(params); err != nil {
			return err
		}
	}
	fmt.Println("")
	return nil
}

// NewPostgresWithDB works just like NewPostgres but with *gorm.DB pointer specified
func NewPostgresWithDB(db *gorm.DB) Postgres { return Postgres{db: db} }

// NewPostgres creates a new PostgreSQL-based storage via GORM and default/loaded DB settings
func NewPostgres() Postgres {
	p := Postgres{}
	p.db = p.DB()
	return p
}

// Postgres implements Keeper for postgres
type Postgres struct {
	db *gorm.DB
	sync.Mutex
}

// AddMultiple entries
func (s Postgres) AddMultiple(entries []models.Param) error {
	sql := fmt.Sprintf(`INSERT INTO %s (created_at) VALUES `, models.Param{}.TableName()) // (created_at, ...)
	params := []interface{}{}
	for _, entry := range entries {
		sql += `(?), ` // (?, ?, ...)
		params = append(params, time.Now() /*, entry. ...*/)
		_ = entry
	}
	sql = strings.TrimSuffix(sql, ", ")

	return s.DB().Exec(sql, params...).Error
}

// Add a new entry
func (s Postgres) Add(entry *models.Param) error {
	return s.DB().Save(entry).Error
}

// Get a model by ID
func (s Postgres) Get(ID uint) (*models.Param, error) {
	var model models.Param
	err := s.DB().First(&model, ID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &model, nil
}

// DeleteAll deletes all rows
func (s Postgres) DeleteAll() error {
	return s.DB().Unscoped().Delete(models.Param{}).Error
}
