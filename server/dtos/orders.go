package dtos

import "time"

type OrderInfo struct {
	OrderId             uint
	OrderName           string
	CustomerCompanyName string
	CustomerName        string
	OrderDate           time.Time
	DeliveredAmount     float64
	TotalAmount         float64
}

type GetOrdersReq struct {
	Filters struct {
		DateRange struct {
			To   uint
			From uint
		}
	}
}
type GetOrdersResp []OrderInfo
