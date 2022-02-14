package service

import "bootcamp/data"

type IPutService interface {
	BalanceInfo(string) (string, error)
}

type PutService struct {
	Data data.IData
}

func (*PutService) BalanceInfo(s string) (string, error) {
	return "hello put service, " + s, nil
}

func NewPutService(data data.IData) IPutService {
	return &PutService{Data: data}
}
