package models

// Param represents configuration entry
type Param struct {
	ID    uint   `gorm:"column:id;primary_key"`
	Type  string `gorm:"column:type"`
	Data  string `gorm:"column:data"`
	Value []byte `json:"column:value"`
}

// TableName for the model
func (Param) TableName() string { return "param" }
