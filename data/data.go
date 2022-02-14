package data

import (
	"encoding/json"
	"strconv"
)

type IData interface {
	GetAllBalanceInfo() (string, error)
	GetBalanceInfo(userName string) (string, error)
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

func (c *Data) GetAllBalanceInfo() (string, error) {
	mapJsonFormat, err := json.Marshal(Balance)
	return string(mapJsonFormat), err
}

func (c *Data) GetBalanceInfo(userName string) (string, error) {
	return strconv.Itoa(Balance[userName]), nil
}

func NewData(initialBalanceAmount, minimumBalanceAmount int) IData {
	return &Data{initialBalanceAmount: initialBalanceAmount,
		minimumBalanceAmount: minimumBalanceAmount}
}
