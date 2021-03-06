package handler

import (
	service "bootcamp/services"
	"fmt"
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

	response := ""
	userName := r.URL.Path

	// ex route: /:username/cem
	// output: cem
	if len(r.URL.Path) > 1 {
		userName = r.URL.Path[len("/:username/"):]
	}

	// validation stage
	// r.URL.Path -> params
	// r.Method -> GET, POST, PUT
	if r.Method == "GET" {
		if userName == "/" {
			response = c.getService.AllBalanceInfo()
		} else {
			// "GET /:username/{userName}" endpoint
			response = c.getService.BalanceInfo(userName)
		}
	} else if r.Method == "PUT" {
		if userName == "/" {
			return
		} else {
			// "PUT /:username/{userName}" endpoint
			c.putService.SetWallet(userName)
		}
	} else if r.Method == "POST" {
		if userName == "/" {
			return
		} else {
			// "POST /:username/{balance}" endpoint
			errType := c.postService.UpdateWallet(userName)
			if errType == 500 {
				response = "PostService: Internal Server Error"
			}
		}
	} else {
		// for now only allowing GET, PUT and POST
		return
	}

	// formatting stage
	//w.Write([]byte(string(response["cem"])))
	fmt.Println(response)
}

func NewHandler(getService service.IGetService, putService service.IPutService, postService service.IPostService) IHandler {
	return &Handler{getService: getService,
		putService:  putService,
		postService: postService}
}
