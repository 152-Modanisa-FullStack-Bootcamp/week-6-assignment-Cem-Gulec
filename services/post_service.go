package service

type IPostService interface {
	BalanceInfo(string) string
}

type PostService struct {
}

func (*PostService) BalanceInfo(s string) string {
	return "hello, " + s
}

func NewPostService() IPostService {
	return &PostService{}
}
