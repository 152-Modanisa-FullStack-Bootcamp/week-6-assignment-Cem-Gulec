package service

import "bootcamp/data"

type IGetService interface {
	BalanceInfo(string) string
}

type GetService struct {
	Data data.IData
}

func (*GetService) BalanceInfo(s string) string {
	return "hello, " + s
}

func NewGetService(data data.IData) IGetService {
	return &GetService{Data: data}
}
