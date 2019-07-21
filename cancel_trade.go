package indodax

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type respCancelOrder struct {
	Success int
	Return  *CancelOrder
	Message string
}

//
// CancelOrder contains a success response from calling a "cancelOrder"
// method.
//
type CancelOrder struct {
	OrderID  int64
	Type     string
	PairName string
	Balance  map[string]float64
}

func (cancelOrder *CancelOrder) UnmarshalJSON(b []byte) (err error) {
	var kv map[string]interface{}

	err = json.Unmarshal(b, &kv)
	if err != nil {
		return err
	}

	for k, v := range kv {
		k = strings.ToLower(k)

		switch k {
		case fieldNameBalance:
			cancelOrder.Balance, err = jsonToMapStringFloat64(v.(map[string]interface{}))
		case fieldNameOrderID:
			cancelOrder.OrderID, err = strconv.ParseInt(fmt.Sprintf("%v", v), 10, 64)
		case fieldNameType:
			cancelOrder.Type = v.(string)
		case fieldNamePairName:
			cancelOrder.PairName = v.(string)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
