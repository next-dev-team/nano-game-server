package model

import "time"

type BaseModel struct {
	CreatedBy *uint
	UpdatedBy *uint
	CreatedAt *time.Time `gorm:"index"`
	UpdatedAt *time.Time `gorm:"index"`
}
