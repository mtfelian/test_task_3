package storage

import (
	"github.com/mtfelian/test_task_3/model1"
)

// Keeper provides access log storage abstraction
type Keeper interface {
	GenerateTestEntries(amount, bulkSize uint) error
	Add(entry *model1.Model) error
	AddMultiple(entries []model1.Model) error
	Get(ID uint) (*model1.Model, error)
	DeleteAll() error
}
