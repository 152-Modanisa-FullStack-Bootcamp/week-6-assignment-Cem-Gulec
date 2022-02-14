package service

import (
	"bootcamp/data"
)

type IGetService interface {
	BalanceInfo(string) (string, error)
}

type GetService struct {
	Data data.IData
}

func (s *GetService) BalanceInfo(userName string) (string, error) {

	// "GET /" endpoint
	if userName == "/" {
		return s.Data.GetAllBalanceInfo()
	} else {
		// "GET /:username" endpoint
		return s.Data.GetBalanceInfo(userName[1:])
	}
}

func NewGetService(data data.IData) IGetService {
	return &GetService{Data: data}
}
