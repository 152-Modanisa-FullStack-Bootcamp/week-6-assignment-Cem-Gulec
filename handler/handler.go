package handler

import (
	service "bootcamp/services"
	"net/http"
)

type IHandler interface {
	Wallet(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	getService  service.IGetService
	putService  service.IPutService
	postService service.IPostService
}

func (c *Handler) Wallet(w http.ResponseWriter, r *http.Request) {

	//var response map[string]int

	// validation stage
	// r.URL.Path -> params
	// r.Method -> GET, POST, PUT
	if r.Method == "GET" {
		//response, _ = c.getService.BalanceInfo(r.URL.Path)
	} else if r.Method == "PUT" {
		return
	} else if r.Method == "POST" {
		return
	} else {
		return
	}

	// formatting stage
	w.Write([]byte("response"))
}

func NewHandler(getService service.IGetService, putService service.IPutService, postService service.IPostService) IHandler {
	return &Handler{getService: getService, putService: putService, postService: postService}
}
