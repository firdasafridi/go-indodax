package indodax

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type respOrderHistory struct {
	Success int
	Return  *respOrders
	Message string
}

type respOrders struct {
	Orders []OrderHistory
}

//
// Order History containt all order book from user
//
type OrderHistory struct {
	ID           int64
	Type         string
	Price        float64
	SubmitTime   time.Time
	FinishTime   time.Time
	Status       string
	OrderAmount  float64
	RemainAmount float64
	AssetName    string
}

func (orderHistory *OrderHistory) UnmarshalJSON(b []byte) (err error) {
	var kv map[string]interface{}

	err = json.Unmarshal(b, &kv)
	if err != nil {
		return err
	}

	for k, v := range kv {
		k = strings.ToLower(k)

		switch k {
		case fieldNameOrderID:
			orderHistory.ID, err = strconv.ParseInt(v.(string), 10, 64)
		case fieldNameType:
			orderHistory.Type = v.(string)
		case fieldNameStatus:
			orderHistory.Status = v.(string)
		case fieldNamePrice:
			orderHistory.Price, err = strconv.ParseFloat(v.(string), 64)
		case fieldNameSubmitTime:
			ts, err := strconv.ParseInt(v.(string), 10, 64)
			if err != nil {
				return err
			}
			orderHistory.SubmitTime = time.Unix(ts, 0)
		case fieldNameFinishTime:
			ts, err := strconv.ParseInt(v.(string), 10, 64)
			if err != nil {
				return err
			}
			orderHistory.FinishTime = time.Unix(ts, 0)
		default:
			keyName := strings.Split(k, "_")
			if len(keyName) < 2 {
				continue
			}
			if keyName[0] == "order" {
				orderHistory.OrderAmount, err = strconv.ParseFloat(v.(string), 64)
			}
			if keyName[0] == "remain" {
				orderHistory.RemainAmount, err = strconv.ParseFloat(v.(string), 64)
			}
			orderHistory.AssetName = keyName[1]

		}
		if err != nil {
			return err
		}
	}
	return nil
}
