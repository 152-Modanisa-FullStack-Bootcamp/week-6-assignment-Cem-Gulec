package main

import (
	"bootcamp/config"
	"bootcamp/data"
	"bootcamp/handler"
	service "bootcamp/services"
	"fmt"
	"net/http"
)

func main() {
	initialBalanceAmount := config.C.InitialBalanceAmount
	minimumBalanceAmount := config.C.MinimumBalanceAmount
	data := data.NewData(initialBalanceAmount, minimumBalanceAmount)

	getService := service.NewGetService(data)
	handler := handler.NewHandler(getService)

	http.HandleFunc("/", handler.Wallet)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
	}
}
