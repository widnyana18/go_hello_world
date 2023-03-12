package crud

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type Order struct {
	OrderId  string    `json: "orderId" example: "1"`
	CustName string    `json: "customerName" example: "Franky Frex"`
	OrderAt  time.Time `json: "orderAt" example: "2019-11-09T21:21:46+00:00"`
	Items    []Item    `json: "items"`
}

type Item struct {
	ItemId string `json: "itemId" example: "A12BT"`
	Desc   string `json: "description" example: "A random Description"`
	Qty    int    `json: "quantity" example: "5"`
}

var orderDt []Order
var prevOrderId = 1

// GetOrders godoc
// @Summary Get Details of All Orders
// @Description Get Details of All Orders
// @Tags Orders
// @Accept  json
// @Produce  json
// @Success 200 {array} Order
// @Router /orders [get]
func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "apllication/json")
	json.NewEncoder(w).Encode(orderDt)
}

// GetOrderById godoc
// @Summary Get a order
// @Description get order with id
// @Tags Order
// @Accept  json
// @Produce  json
// @Success 200 {array} Order
// @Router /order/:orderId [get]
func GetOrderById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "apllication/json")

	for _, order := range orderDt {
		if order.OrderId == r.URL.Path {
			json.NewEncoder(w).Encode(order)
		}
		break
	}
}

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order with the input paylod
// @Tags Orders
// @Accept  json
// @Produce  json
// @Param create-order body Order true "Create Order"
// @Success 200 {array} Order
// @Router /create-order [post]
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "apllication/json")

	var order Order
	json.NewDecoder(r.Body).Decode(&order)
	prevOrderId++
	order.OrderId = strconv.Itoa(prevOrderId)
	orderDt = append(orderDt, order)

	json.NewEncoder(w).Encode(order)
}

// UpdateOrder godoc
// @Summary Update order
// @Description update data order
// @Tags Orders
// @Accept  json
// @Produce  json
// @Param updateOrder body Order true "update order"
// @Success 200 {array} Order
// @Router /order/:orderId/update [put]
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "apllication/json")

	id := r.URL.Path
	var newOrder Order

	json.NewDecoder(r.Body).Decode(&newOrder)

	for _, order := range orderDt {
		if order.OrderId == id {
			order = newOrder
		}
		break
	}

	json.NewEncoder(w).Encode(newOrder)
}

// DeleteOrder godoc
// @Summary delete order
// @Description Get Details of All Orders
// @Tags Orders
// @Accept  json
// @Produce  json
// @Success 200 {array} Order
// @Router order/:orderId [delete]
func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "apllication/json")

	for _, order := range orderDt {
		if order.OrderId == r.URL.Path {
			order = Order{}
			json.NewEncoder(w).Encode(order)
		}
		break
	}
}
