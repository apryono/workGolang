package main

import (
	"fmt"
	"net/http"
	detail "tugas/GoResto/detail/handler"
	menu "tugas/GoResto/menu/handler"
	order "tugas/GoResto/order/handler"
	table "tugas/GoResto/table/handler"

	"github.com/gorilla/mux"
)

func main() {
	port := "8080"

	r := mux.NewRouter().StrictSlash(true)

	menu.CreateMenuHandler(r)
	table.CreateTableHandler(r)
	order.CreateOrderHandler(r)
	detail.CreateDetailHandler(r)

	fmt.Println("Starting Web Server at port : " + port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		fmt.Printf("[main.ListenAndServe] Error when serving with error : %v \n", err.Error())
	}

}
