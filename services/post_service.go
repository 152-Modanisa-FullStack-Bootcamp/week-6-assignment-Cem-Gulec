package service

import "bootcamp/data"

type IPostService interface {
	BalanceInfo(string) string
}

type PostService struct {
	Data data.IData
}

func (*PostService) BalanceInfo(s string) string {
	return "hello, " + s
}

func NewPostService(data data.IData) IPostService {
	return &PostService{Data: data}
}
