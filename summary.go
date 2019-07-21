package indodax

import (
	"encoding/json"
	"fmt"
	"strings"
)

//
// Summary containts all tickers, prices24h, and prices 7d status of all pairs.
//
type Summary struct {
	Tickers   map[string]*Ticker
	Prices24h map[string]float64
	Prices7d  map[string]float64
}

func (sum *Summary) UnmarshalJSON(b []byte) (err error) {
	var kv map[string]map[string]interface{}

	err = json.Unmarshal(b, &kv)
	if err != nil {
		return err
	}

	tickers := make(map[string]*Ticker, len(kv[fieldNameTickers]))
	prices24h := make(map[string]float64, len(kv[fieldNamePrices24h]))
	prices7d := make(map[string]float64, len(kv[fieldNamePrices7d]))

	for k, v := range kv {
		switch k {
		case fieldNameTickers:
			tickers, err = mapStringToTickers(v)
		case fieldNamePrices24h:
			prices24h, err = jsonToMapStringFloat64(v)
		case fieldNamePrices7d:
			prices7d, err = jsonToMapStringFloat64(v)
		}
		if err != nil {
			return err
		}
	}
	sum.Tickers = tickers
	sum.Prices24h = prices24h
	sum.Prices7d = prices7d

	return nil
}

func mapStringToTickers(tickerInterface map[string]interface{}) (sum map[string]*Ticker, err error) {

	sum = make(map[string]*Ticker, len(tickerInterface))

	for k, v := range tickerInterface {
		ticker, err := mapInterfaceToTicker(k, v)
		if err != nil {
			return nil, err
		}
		sum[k] = ticker
	}

	return sum, nil
}

func mapInterfaceToTicker(pairName string, tkrInterface interface{}) (tkr *Ticker, err error) {

	b, err := json.Marshal(tkrInterface)
	if err != nil {
		return nil, err
	}

	var kv map[string]interface{}
	err = json.Unmarshal(b, &kv)
	if err != nil {
		return nil, err
	}

	var tkrString ticker
	for k, v := range kv {
		k = strings.ToLower(k)

		valStr, ok := v.(string)
		if !ok {
			valStr = fmt.Sprintf("%v", v)
		}

		switch k {
		case fieldNameHigh:
			tkrString.High = valStr
		case fieldNameLow:
			tkrString.Low = valStr
		case fieldNameAmount:
			tkrString.Last = valStr
		case fieldNameBuy:
			tkrString.Buy = valStr
		case fieldNameSell:
			tkrString.Sell = valStr
		default:
			if !strings.HasPrefix(k, "vol_") {
				continue
			}
			volName := strings.Split(k, "_")
			if len(volName) != 2 {
				continue
			}
			if tkrString.volumes == nil {
				tkrString.volumes = make(map[string]string)
			}
			tkrString.volumes[volName[1]] = valStr
		}
	}

	tkr, err = tkrString.toTicker(pairName)
	if err != nil {
		return nil, err
	}
	return tkr, nil
}
