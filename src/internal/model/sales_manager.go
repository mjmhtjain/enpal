package model

import (
	"github.com/lib/pq"
)

type SalesManager struct {
	ID              uint           `gorm:"primaryKey"`
	Name            string         `gorm:"column:name;not null"`
	Languages       pq.StringArray `gorm:"column:languages;type:varchar[]"`
	Products        pq.StringArray `gorm:"column:products;type:varchar[]"`
	CustomerRatings pq.StringArray `gorm:"column:customer_ratings;type:varchar[]"`
	Slots           []Slot         `gorm:"foreignKey:SalesManagerID"`
}

// TableName specifies the table name for GORM
func (SalesManager) TableName() string {
	return "sales_managers"
}
