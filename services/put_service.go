package service

type IPutService interface {
	BalanceInfo(string) string
}

type PutService struct {
}

func (*PutService) BalanceInfo(s string) string {
	return "hello, " + s
}

func NewPutService() IPutService {
	return &PutService{}
}
