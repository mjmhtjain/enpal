package model

type SalesManager struct {
	ID              uint   `gorm:"primaryKey"`
	Name            string `gorm:"column:name;not null"`
	Languages       string `gorm:"column:languages"`
	Products        string `gorm:"column:products"`
	CustomerRatings string `gorm:"column:customer_ratings"`
	Slots           []Slot `gorm:"foreignKey:SalesManagerID"`
}

// TableName specifies the table name for GORM
func (SalesManager) TableName() string {
	return "sales_managers"
}
