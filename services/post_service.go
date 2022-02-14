package service

import "bootcamp/data"

type IPostService interface {
	UpdateWallet(string) int
}

type PostService struct {
	Data data.IData
}

func (p *PostService) UpdateWallet(userName string) int {
	return p.Data.PostWalletInfo(userName)
}

func NewPostService(data data.IData) IPostService {
	return &PostService{Data: data}
}
