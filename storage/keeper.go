package storage

import (
	"github.com/mtfelian/test_task_3/models"
)

// Keeper provides access log storage abstraction
type Keeper interface {
	GenerateTestEntries(amount, bulkSize uint) error
	Add(entry *models.Param) error
	AddMultiple(entries []models.Param) error
	Get(ID uint) (*models.Param, error)
	DeleteAll() error
}
