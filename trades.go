package indodax

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Trade struct {
	ID     int64
	Type   string
	Date   time.Time
	Amount float64
	Price  float64
}

func (trade *Trade) UnmarshalJSON(b []byte) (err error) {
	var kv map[string]string

	err = json.Unmarshal(b, &kv)
	if err != nil {
		return err
	}

	for k, v := range kv {
		k = strings.ToLower(k)

		switch k {
		case fieldNameTID:
			trade.ID, err = strconv.ParseInt(v, 10, 64)
		case fieldNameType:
			trade.Type = v
		case fieldNameDate:
			ts, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return err
			}
			trade.Date = time.Unix(ts, 0)
		case fieldNameAmount:
			trade.Amount, err = strconv.ParseFloat(v, 64)
		case fieldNamePrice:
			trade.Price, err = strconv.ParseFloat(v, 64)
		}

		if err != nil {
			fmt.Println(k)
			return err
		}
	}

	return nil
}
