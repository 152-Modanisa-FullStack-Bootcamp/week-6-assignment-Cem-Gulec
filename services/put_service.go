package service

import "bootcamp/data"

type IPutService interface {
	BalanceInfo(string) string
}

type PutService struct {
	Data data.IData
}

func (*PutService) BalanceInfo(s string) string {
	return "hello, " + s
}

func NewPutService(data data.IData) IPutService {
	return &PutService{Data: data}
}
