package indodax

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type respGetOrders struct {
	Success int
	Return  respGetOrder
	Message string
}

type respGetOrder struct {
	Order *GetOrder
}

//
// Get Order containt a status from order placed of user
//
type GetOrder struct {
	OrderID      int64
	Price        float64
	Type         string
	OrderAmount  float64
	RemainAmount float64
	SubmitTime   time.Time
	FinishTime   time.Time
	Status       string
	AssetName    string
}

func (getOrder *GetOrder) UnmarshalJSON(b []byte) (err error) {
	var kv map[string]string

	err = json.Unmarshal(b, &kv)
	if err != nil {
		return err
	}

	for k, v := range kv {
		k = strings.ToLower(k)

		switch k {
		case fieldNameOrderID:
			getOrder.OrderID, err = strconv.ParseInt(v, 10, 64)
		case fieldNameType:
			getOrder.Type = v
		case fieldNameStatus:
			getOrder.Status = v
		case fieldNameSubmitTime:
			ts, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return err
			}
			getOrder.SubmitTime = time.Unix(ts, 0)
		case fieldNameFinishTime:
			if v == "0" {
				continue
			}
			ts, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return err
			}
			getOrder.FinishTime = time.Unix(ts, 0)
		case fieldNamePrice:
			getOrder.Price, err = strconv.ParseFloat(v, 64)
		default:
			keyName := strings.Split(k, "_")
			if len(keyName) < 2 {
				continue
			}

			switch keyName[0] {
			case fieldNameRemain:
				getOrder.RemainAmount, err = strconv.ParseFloat(v, 64)
			case fieldNameOrder:
				getOrder.OrderAmount, err = strconv.ParseFloat(v, 64)
			}
			getOrder.AssetName = keyName[1]
		}

		if err != nil {
			return err
		}
	}
	return nil
}
