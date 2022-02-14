package service

import (
	"bootcamp/data"
	"encoding/json"
	"strconv"
)

type IGetService interface {
	AllBalanceInfo() string
	BalanceInfo(userName string) string
}

type GetService struct {
	Data data.IData
}

func (s *GetService) AllBalanceInfo() string {
	m := s.Data.GetBalanceInfo()
	mapJsonFormat, _ := json.Marshal(m)

	return string(mapJsonFormat)
}

func (s *GetService) BalanceInfo(userName string) string {
	m := s.Data.GetBalanceInfo()

	return strconv.Itoa(m[userName[1:]])
}

func NewGetService(data data.IData) IGetService {
	return &GetService{Data: data}
}
