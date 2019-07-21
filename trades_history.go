package indodax

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type respTradeHistory struct {
	Success int
	Return  *respTradeHist
	Message string
}

type respTradeHist struct {
	Trades []TradeHistory
}

//
// Trade History containt trade id, order id, type of trade match(buy/sell), AssetName, amount, price, and fee.
//
type TradeHistory struct {
	TradeID   int64
	OrderID   int64
	Type      string
	AssetName string
	Amount    float64
	Price     float64
	Fee       float64
	TradeTime time.Time
}

func (tradeHistory *TradeHistory) UnmarshalJSON(b []byte) (err error) {
	var kv map[string]interface{}

	err = json.Unmarshal(b, &kv)
	if err != nil {
		return err
	}
	for k, v := range kv {
		k = strings.ToLower(k)
		switch k {
		case fieldNameOrderID:
			tradeHistory.OrderID, err = strconv.ParseInt(v.(string), 10, 64)
		case fieldNameTradeID:
			tradeHistory.TradeID, err = strconv.ParseInt(v.(string), 10, 64)
		case fieldNameType:
			tradeHistory.Type = v.(string)
		case fieldNameTradeTime:
			ts, err := strconv.ParseInt(v.(string), 10, 64)
			if err != nil {
				return err
			}
			tradeHistory.TradeTime = time.Unix(ts, 0)
		case fieldNamePrice:
			tradeHistory.Price, err = strconv.ParseFloat(v.(string), 64)
		case fieldNameFee:
			tradeHistory.Fee, err = strconv.ParseFloat(v.(string), 64)
		default:
			tradeHistory.Amount, err = strconv.ParseFloat(v.(string), 64)
			tradeHistory.AssetName = k
		}

		if err != nil {
			return err
		}
	}
	return nil
}
