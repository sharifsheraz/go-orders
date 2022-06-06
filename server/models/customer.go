package models

import "github.com/lib/pq"

type Customer struct {
	ID          string `json:"user_id" gorm:"primaryKey"`
	Username    string `json:"login" gorm:"not null"`
	Password    string `gorm:"not null"`
	Name        string `gorm:"not null"`
	CompanyID   uint   `gorm:"not null"`
	Company     Company
	CreditCards pq.StringArray `gorm:"type:text[];not null"`
}
