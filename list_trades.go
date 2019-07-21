package indodax

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//
// List Trade containt all match order from user
//
type ListTrade struct {
	ID     int64
	Type   string
	Date   time.Time
	Amount float64
	Price  float64
}

func (listTrade *ListTrade) UnmarshalJSON(b []byte) (err error) {
	var kv map[string]string

	err = json.Unmarshal(b, &kv)
	if err != nil {
		return err
	}

	for k, v := range kv {
		k = strings.ToLower(k)

		switch k {
		case fieldNameTID:
			listTrade.ID, err = strconv.ParseInt(v, 10, 64)
		case fieldNameType:
			listTrade.Type = v
		case fieldNameDate:
			ts, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return err
			}
			listTrade.Date = time.Unix(ts, 0)
		case fieldNameAmount:
			listTrade.Amount, err = strconv.ParseFloat(v, 64)
		case fieldNamePrice:
			listTrade.Price, err = strconv.ParseFloat(v, 64)
		}

		if err != nil {
			fmt.Println(k)
			return err
		}
	}

	return nil
}
