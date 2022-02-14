package service

import "bootcamp/data"

type IPostService interface {
	BalanceInfo(string) (string, error)
}

type PostService struct {
	Data data.IData
}

func (*PostService) BalanceInfo(s string) (string, error) {
	return "hello post service, " + s, nil
}

func NewPostService(data data.IData) IPostService {
	return &PostService{Data: data}
}
