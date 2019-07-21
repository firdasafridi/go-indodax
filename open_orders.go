package indodax

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

//
// Open Orders containt all order book from user
//
type OpenOrders struct {
	ID           int64
	SubmitTime   time.Time
	Price        float64
	OrderAmount  float64
	RemainAmount float64
	Type         string
	AssetName    string
}

type responseAllOpenOrders struct {
	Success int
	Return  respAllOrder
	Message string
}

type respAllOrder struct {
	Orders map[string][]OpenOrders
}

type responseOpenOrders struct {
	Success int
	Return  respOrder
	Message string
}

type respOrder struct {
	Orders []OpenOrders
}

func (openOrders *OpenOrders) UnmarshalJSON(b []byte) (err error) {
	var kv map[string]interface{}

	err = json.Unmarshal(b, &kv)
	if err != nil {
		return err
	}

	for k, v := range kv {
		k = strings.ToLower(k)

		switch k {
		case fieldNameOrderID:
			openOrders.ID, err = strconv.ParseInt(v.(string), 10, 64)
		case fieldNameType:
			openOrders.Type = v.(string)
		case fieldNameSubmitTime:
			ts, err := strconv.ParseInt(v.(string), 10, 64)
			if err != nil {
				return err
			}
			openOrders.SubmitTime = time.Unix(ts, 0)
		case fieldNamePrice:
			price := v.(string)
			openOrders.Price, err = strconv.ParseFloat(price, 64)
		default:
			keyName := strings.Split(k, "_")
			if len(keyName) < 2 {
				continue
			}

			switch keyName[0] {
			case fieldNameRemain:
				openOrders.RemainAmount, err = strconv.ParseFloat(v.(string), 64)
			case fieldNameOrder:
				openOrders.OrderAmount, err = strconv.ParseFloat(v.(string), 64)
			}
			openOrders.AssetName = keyName[1]
		}
		if err != nil {
			return err
		}
	}
	return nil
}
