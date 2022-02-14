package handler

import (
	service "bootcamp/services"
	"net/http"
)

type IHandler interface {
	Wallet(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	service service.IGetService
}

func (c *Handler) Wallet(w http.ResponseWriter, r *http.Request) {

	/*// validation stage
	if r.Method != "GET" {
		return
	}*/

	result := c.service.BalanceInfo(r.URL.Path)

	// formatting stage
	w.Write([]byte(result))
}

func NewHandler(service service.IGetService) IHandler {
	return &Handler{service: service}
}
