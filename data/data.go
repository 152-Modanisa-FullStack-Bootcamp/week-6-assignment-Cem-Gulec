package data

type IData interface {
	GetBalanceInfo() map[string]int
	PutWallet(string)
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

func (c *Data) PutWallet(userName string) {
	// if element with given username exists
	if _, ok := Balance[userName]; ok {
		return
	} else {
		Balance[userName] = c.initialBalanceAmount
	}
}

func NewData(initialBalanceAmount, minimumBalanceAmount int) IData {
	return &Data{initialBalanceAmount: initialBalanceAmount,
		minimumBalanceAmount: minimumBalanceAmount}
}
