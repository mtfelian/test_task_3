package model1

import "time"

// Model of the model1
type Model struct {
	ID        uint      `gorm:"column:id;primary_key"`
	CreatedAt time.Time `gorm:"column:created_at"`
	// something
}

// TableName for the model
func (Model) TableName() string { return "model1" }
