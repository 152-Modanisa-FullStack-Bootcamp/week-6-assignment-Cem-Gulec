package main

import (
	"bootcamp/config"
	"bootcamp/handler"
	service "bootcamp/services"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println(config.C.InitialBalanceAmount)
	service := service.NewGetService()
	handler := handler.NewHandler(service)

	http.HandleFunc("/", handler.Wallet)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
	}
}
