package model

import "time"

type Slot struct {
	ID             uint         `gorm:"primaryKey"`
	StartDate      time.Time    `gorm:"column:start_date;not null"`
	EndDate        time.Time    `gorm:"column:end_date;not null"`
	Booked         bool         `gorm:"column:booked;default:false"`
	SalesManagerID uint         `gorm:"column:sales_manager_id"`
	SalesManager   SalesManager `gorm:"foreignKey:SalesManagerID"`
}

// TableName specifies the table name for GORM
func (Slot) TableName() string {
	return "slots"
}
