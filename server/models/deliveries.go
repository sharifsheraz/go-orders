// id,order_item_id,delivered_quantity
package models

type Delivery struct {
	ID                uint `gorm:"primaryKey"`
	OrderItemID       uint `gorm:"not null"`
	DeliveredQuantity uint `gorm:"not null"`
}
