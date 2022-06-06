package main

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/sharifsheraz/go-orders/db"
	"github.com/sharifsheraz/go-orders/models"
	"github.com/sharifsheraz/go-orders/util"
	"gorm.io/gorm"
)

func main() {
	conn := db.Init()
	SeedData(conn)
}
func SeedData(conn *gorm.DB) {
	SeedCompanies(conn)
	SeedCustomers(conn)
	SeedOrders(conn)
	SeedOrderItems(conn)
	SeedDeliveries(conn)
}
func SeedCompanies(db *gorm.DB) {
	var companies []models.Company
	records := util.ReadCSV("test_data/customer_companies.csv")

	for _, record := range records {
		id, err := strconv.ParseUint(record[0], 10, 32)
		if err != nil {
			panic(err)
		}
		company := models.Company{
			ID:   uint(id),
			Name: record[1],
		}
		companies = append(companies, company)

	}
	if result := db.Create(companies); result.Error != nil {
		panic(result.Error)
	}
}

func SeedCustomers(db *gorm.DB) {
	var customers []models.Customer
	records := util.ReadCSV("test_data/customers.csv")
	for _, record := range records {
		companyID, err := strconv.ParseUint(record[4], 10, 32)
		if err != nil {
			panic(err)
		}

		var creditCards []string
		err = json.Unmarshal([]byte(record[5]), &creditCards)
		if err != nil {
			log.Fatal(err)
		}

		customer := models.Customer{
			ID:          record[0],
			Username:    record[1],
			Password:    record[2],
			Name:        record[3],
			CompanyID:   uint(companyID),
			CreditCards: creditCards,
		}

		customers = append(customers, customer)

	}
	if result := db.Create(customers); result.Error != nil {
		panic(result.Error)
	}
}

func SeedOrders(db *gorm.DB) {
	var orders []models.Order
	records := util.ReadCSV("test_data/orders.csv")

	for _, record := range records {
		id, err := strconv.ParseUint(record[0], 10, 32)
		if err != nil {
			panic(err)
		}
		createdAt, err := time.Parse(time.RFC3339, record[1])

		if err != nil {
			panic(err)
		}

		order := models.Order{
			ID:         uint(id),
			CreatedAt:  createdAt,
			OrderName:  record[2],
			CustomerID: record[3],
		}
		orders = append(orders, order)

	}
	if result := db.Create(orders); result.Error != nil {
		panic(result.Error)
	}
}
func SeedOrderItems(db *gorm.DB) {
	var orderItems []models.OrderItem
	records := util.ReadCSV("test_data/order_items.csv")

	for _, record := range records {
		orderItemID, err := strconv.ParseUint(record[0], 10, 32)
		if err != nil {
			panic(err)
		}
		orderID, err := strconv.ParseUint(record[1], 10, 32)
		if err != nil {
			panic(err)
		}
		price, err := strconv.ParseFloat(record[2], 32)
		if err != nil {
			price = 0
		}
		quantity, err := strconv.ParseUint(record[3], 10, 32)
		if err != nil {
			panic(err)
		}
		orderItem := models.OrderItem{
			ID:           uint(orderItemID),
			OrderID:      uint(orderID),
			PricePerUnit: float64(price),
			Quantity:     uint(quantity),
			Product:      record[4],
		}
		orderItems = append(orderItems, orderItem)
	}
	if result := db.Create(orderItems); result.Error != nil {
		panic(result.Error)
	}
}
func SeedDeliveries(db *gorm.DB) {
	var deliveries []models.Delivery
	records := util.ReadCSV("test_data/deliveries.csv")

	for _, record := range records {
		deliveryID, err := strconv.ParseUint(record[0], 10, 32)
		if err != nil {
			panic(err)
		}

		orderItemID, err := strconv.ParseUint(record[1], 10, 32)
		if err != nil {
			panic(err)
		}

		dQuantity, err := strconv.ParseUint(record[2], 10, 32)
		if err != nil {
			panic(err)
		}

		orderItem := models.Delivery{
			ID:                uint(deliveryID),
			OrderItemID:       uint(orderItemID),
			DeliveredQuantity: uint(dQuantity),
		}
		deliveries = append(deliveries, orderItem)
	}
	if result := db.Create(deliveries); result.Error != nil {
		panic(result.Error)
	}
}
