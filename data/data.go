package data

type IData interface {
	GetBalanceInfo() map[string]int
}

type Data struct {
	initialBalanceAmount int
	minimumBalanceAmount int
}

// key: username
// value: balance
var Balance = map[string]int{
	"cem":   123,
	"rıfkı": 42,
}

func (c *Data) GetBalanceInfo() map[string]int {
	return Balance
}

func NewData(initialBalanceAmount, minimumBalanceAmount int) IData {
	return &Data{initialBalanceAmount: initialBalanceAmount,
		minimumBalanceAmount: minimumBalanceAmount}
}
