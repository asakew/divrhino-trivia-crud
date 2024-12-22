package models

import "gorm.io/gorm"

type Fact struct {
	gorm.Model
	ID       string `gorm:"primaryKey" json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func (b *Fact) TableName() string {
	return "facts"
}
