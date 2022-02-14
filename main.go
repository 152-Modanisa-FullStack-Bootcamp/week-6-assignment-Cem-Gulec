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
	initialBalanceAmount := config.Get().InitialBalanceAmount
	minimumBalanceAmount := config.Get().MinimumBalanceAmount
	data := data.NewData(initialBalanceAmount, minimumBalanceAmount)

	getService := service.NewGetService(data)
	putService := service.NewPutService(data)
	postService := service.NewPostService(data)
	handler := handler.NewHandler(getService, putService, postService)

	http.HandleFunc("/", handler.Wallet)
	http.HandleFunc("/:username/", handler.Wallet)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
	}
}
