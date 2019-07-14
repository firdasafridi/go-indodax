package indodax

import (
	"encoding/json"
	"fmt"
	"strconv"
)

//
// Order contains the number of amount asset and price of open order
//
type Order struct {
	Amount float64
	Price  float64
}

func (order *Order) UnmarshalJSON(b []byte) (err error) {
	var val []interface{}

	err = json.Unmarshal(b, &val)
	if err != nil {
		return err
	}

	if len(val) != 2 {
		return fmt.Errorf("order: UnmarshalJSON: invalid length of order")
	}

	order.Price, err = strconv.ParseFloat(fmt.Sprintf("%v", val[0]), 64)
	if err != nil {
		return err
	}

	order.Amount, err = strconv.ParseFloat(fmt.Sprintf("%v", val[1]), 64)
	if err != nil {
		return err
	}

	return nil
}
