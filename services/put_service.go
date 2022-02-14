package service

import "bootcamp/data"

type IPutService interface {
	SetWallet(userName string)
}

type PutService struct {
	Data data.IData
}

func (p *PutService) SetWallet(userName string) {
	p.Data.PutWallet(userName)
}

func NewPutService(data data.IData) IPutService {
	return &PutService{Data: data}
}
