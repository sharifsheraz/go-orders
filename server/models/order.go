package models

import (
	"math"
	"time"
)

type Order struct {
	ID         uint `gorm:"primaryKey"`
	CreatedAt  time.Time
	OrderName  string `gorm:"not null"`
	CustomerID string `gorm:"not null"`
	Customer   Customer
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
}

func (o Order) GetOrderAmounts() (float64, float64) {
	tAmount := 0.0
	dAmount := 0.0
	for _, orderItem := range o.OrderItems {
		tAmount += math.Round(float64(orderItem.Quantity)*orderItem.PricePerUnit*100) / 100
		dAmount += orderItem.GetDeliveredAmount()
	}
	return tAmount, dAmount
}
