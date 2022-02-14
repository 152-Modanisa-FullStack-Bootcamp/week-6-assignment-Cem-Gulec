package service

import (
	"bootcamp/data"
	"fmt"
)

type IGetService interface {
	BalanceInfo(string) (map[string]int, error)
}

type GetService struct {
	Data data.IData
}

func (s *GetService) BalanceInfo(userName string) (map[string]int, error) {

	// "GET /" endpoint
	if userName == "/" {
		//return s.Data.GetAllBalanceInfo()
		return nil, nil
	} else {
		// "GET /:username" endpoint
		fmt.Println("olmadı")
		return nil, nil
	}
}

func NewGetService(data data.IData) IGetService {
	return &GetService{Data: data}
}
