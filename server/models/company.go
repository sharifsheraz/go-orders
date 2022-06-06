package models

type Company struct {
	ID   uint   `json:"company_id" gorm:"primaryKey"`
	Name string `json:"company_name" gorm:"not null"`
}
