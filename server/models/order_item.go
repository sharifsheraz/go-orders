package models

import "math"

type OrderItem struct {
	ID           uint    `gorm:"primaryKey"`
	OrderID      uint    `gorm:"not null"`
	PricePerUnit float64 `gorm:"not null"`
	Quantity     uint    `gorm:"not null"`
	Product      string
	Deliveries   []Delivery
}

func (oItem OrderItem) GetDeliveredAmount() float64 {
	dAmount := 0.0
	for _, delivery := range oItem.Deliveries {
		dAmount += math.Round(float64(delivery.DeliveredQuantity)*oItem.PricePerUnit*100) / 100
	}
	return dAmount
}
