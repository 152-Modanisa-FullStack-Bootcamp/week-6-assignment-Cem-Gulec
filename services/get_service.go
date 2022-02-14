package service

type IGetService interface {
	BalanceInfo(string) string
}

type GetService struct {
}

func (*GetService) BalanceInfo(s string) string {
	return "hello, " + s
}

func NewGetService() IGetService {
	return &GetService{}
}
