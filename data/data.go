package data

import (
	"encoding/json"
	"fmt"
)

type IData interface {
	GetAllBalanceInfo() (string, error)
}

type Data struct {
	initialBalanceAmount int
	minimumBalanceAmount int
}

type User struct {
	UserName string
	Balance  int
}

// key: username
// value: balance
Balance := map[string]int{}
Balance := make([]User, N)

func (c *Data) GetAllBalanceInfo() (string, error) {

	Balance["cem"] = 123
	Balance["rıfkı"] = 45
	mapJsonFormat, err := json.Marshal(Balance)

	if err != nil {
		panic(err)
	}

	r := &model.DataResponse{}
	err = json.Unmarshal(mapJsonFormat, r)

	fmt.Println(r)

	return *r, err
}

func NewData(initialBalanceAmount, minimumBalanceAmount int) IData {
	return &Data{initialBalanceAmount: initialBalanceAmount,
		minimumBalanceAmount: minimumBalanceAmount}
}
