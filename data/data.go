package data

import (
	"net/http"
)

type IData interface {
	GetBalanceInfo() map[string]int
	PutWallet(string)
	PostWalletInfo(string) int
}

type Data struct {
	initialBalanceAmount int
	minimumBalanceAmount int
	balance              int
}

// key: username
// value: balance
var Balance = map[string]int{
	"cem":   0,
	"rıfkı": 42,
}

func (*Data) GetBalanceInfo() map[string]int {
	return Balance
}

func (d *Data) PutWallet(userName string) {
	// if element with given username exists
	if _, ok := Balance[userName]; ok {
		return
	} else {
		Balance[userName] = d.initialBalanceAmount
	}
}

func (d *Data) PostWalletInfo(userName string) int {
	if d.balance > 0 {
		Balance[userName] += d.balance
	} else {
		if (Balance[userName] + d.balance) >= d.minimumBalanceAmount {
			Balance[userName] += d.balance
		} else {
			return http.StatusInternalServerError
		}
	}
	return http.StatusOK
}

func NewData(initialBalanceAmount, minimumBalanceAmount, balance int) IData {
	return &Data{initialBalanceAmount: initialBalanceAmount,
		minimumBalanceAmount: minimumBalanceAmount,
		balance:              balance}
}
