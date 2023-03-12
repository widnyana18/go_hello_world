package main

import (
	"go_swagger/crud"
	"log"
	"net/http"
	_ "swaggo-ordes-api/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Orders API
// @version 1.0
// @description This is a sample server for managing orders.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email draco@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("hello World")
	})

	router.HandleFunc("/orders", crud.GetOrders).Methods("GET")
	router.HandleFunc("/order/{orderId}", crud.GetOrderById).Methods("GET")
	router.HandleFunc("/create-order", crud.CreateOrder).Methods("POST")
	router.HandleFunc("/order/{orderId}/update", crud.UpdateOrder).Methods("PUT")
	router.HandleFunc("/order/{orderId}/delete", crud.DeleteOrder).Methods("DELETE")

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	log.Println("connecting server to 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
