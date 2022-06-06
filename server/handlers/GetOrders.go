package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/sharifsheraz/go-orders/dtos"
	"github.com/sharifsheraz/go-orders/models"
	"github.com/sharifsheraz/go-orders/util"
	"gorm.io/gorm/clause"
)

const ClientLocName = "Australia/Melbourne"

func (h handler) GetOrders(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	query := strings.ToLower(params.Get("q"))
	var body dtos.GetOrdersReq
	json.NewDecoder(r.Body).Decode(&body)
	dtRange := &body.Filters.DateRange

	paginated := h.DB.Scopes(util.Paginate(r))
	withAssociations := paginated.Preload(clause.Associations).Preload("Customer." + clause.Associations).Preload("OrderItems." + clause.Associations)
	withFilter := withAssociations

	if dtRange.To != 0 && dtRange.From != 0 {
		to, toErr := util.MillisInLocToTime(dtRange.To, ClientLocName)
		from, fromErr := util.MillisInLocToTime(dtRange.From, ClientLocName)
		if toErr != nil || fromErr != nil {
			log.Print("Error::", toErr, fromErr)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something bad happened!"))
			return
		}
		withFilter = withFilter.Where("created_at BETWEEN ? AND ?", to, from)
	}

	if query != "" {
		var orderIdsMatchingProduct []string
		h.DB.Model(&models.OrderItem{}).Where("lower(product) LIKE ?", "%"+query+"%").Pluck("order_id", &orderIdsMatchingProduct)

		withFilter = withFilter.Where("lower(order_name) LIKE ? OR id IN ? ", "%"+query+"%", orderIdsMatchingProduct)
	}

	var orders []models.Order
	if result := withFilter.Order("created_at DESC").Find(&orders); result.Error != nil {
		log.Print("Error::", result.Error)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	respBody := make(dtos.GetOrdersResp, 0)
	for _, order := range orders {
		orderDAmount, orderTAmount := order.GetOrderAmounts()

		respBody = append(respBody, dtos.OrderInfo{
			OrderId:             order.ID,
			OrderName:           order.OrderName,
			CustomerName:        order.Customer.Name,
			CustomerCompanyName: order.Customer.Company.Name,
			OrderDate:           order.CreatedAt,
			DeliveredAmount:     orderDAmount,
			TotalAmount:         orderTAmount,
		})
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respBody)
}
