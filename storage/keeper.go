package storage

import (
	"github.com/mtfelian/test_task_3/models"
)

// Keeper provides storage abstraction
type Keeper interface {
	Get(_type, data string) (*models.Param, error)
	Close() error // just for testing purposes
}
