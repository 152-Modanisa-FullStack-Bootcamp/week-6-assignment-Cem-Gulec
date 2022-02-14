package data

type IData interface {
	GetBalanceInfo() (int, error)
}

type Data struct {
	initialBalanceAmount int
	minimumBalanceAmount int
}

// key: username
// value: balance
var Balance = map[string]int{}

func (c *Data) GetBalanceInfo() (int, error) {

	return Balance["cem"], nil
}

func NewData(initialBalanceAmount, minimumBalanceAmount int) IData {
	return &Data{initialBalanceAmount: initialBalanceAmount,
		minimumBalanceAmount: minimumBalanceAmount}
}
