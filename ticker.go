package indodax

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

//
// Ticker containts High price 24h, Low price24h, Volume asset Volume Base, Last price, Open  buy, and Open Sell
//
type Ticker struct {
	PairName    string
	High        float64
	Low         float64
	AssetVolume float64
	BaseVolume  float64
	Last        float64
	Buy         float64
	Sell        float64
}

type ticker struct {
	Name    string
	High    string
	Low     string
	volumes map[string]string
	Last    string
	Buy     string
	Sell    string
}

type tickerResponse struct {
	Ticker *ticker
}

func (tkr *tickerResponse) toTicker(pairName string) (ticker *Ticker, err error) {

	ticker, err = tkr.Ticker.toTicker(pairName)
	if err != nil {
		return nil, err
	}

	return ticker, nil
}

func (tkr *ticker) toTicker(pairName string) (ticker *Ticker, err error) {

	high, _ := strconv.ParseFloat(tkr.High, 64)

	low, _ := strconv.ParseFloat(tkr.Low, 64)

	var assetVolume, baseVolume float64 = 0, 0
	volName := strings.Split(pairName, "_")
	if len(tkr.volumes) < 2 {
		return ticker, err
	}
	if v, ok := tkr.volumes[volName[0]]; ok {
		assetVolume, _ = strconv.ParseFloat(v, 64)
	}
	if v, ok := tkr.volumes[volName[1]]; ok {
		baseVolume, _ = strconv.ParseFloat(v, 64)
	}

	last, _ := strconv.ParseFloat(tkr.Last, 64)

	buy, _ := strconv.ParseFloat(tkr.Buy, 64)

	sell, _ := strconv.ParseFloat(tkr.Sell, 64)

	ticker = &Ticker{
		PairName:    pairName,
		High:        high,
		Low:         low,
		AssetVolume: assetVolume,
		BaseVolume:  baseVolume,
		Last:        last,
		Buy:         buy,
		Sell:        sell,
	}

	return ticker, nil
}

func (tkr *ticker) UnmarshalJSON(b []byte) (err error) {
	var kv map[string]interface{}

	err = json.Unmarshal(b, &kv)
	if err != nil {
		return err
	}

	for k, v := range kv {
		k = strings.ToLower(k)

		valStr, ok := v.(string)
		if !ok {
			valStr = fmt.Sprintf("%v", v)
		}

		switch k {
		case fieldNameHigh:
			tkr.High = valStr
		case fieldNameLow:
			tkr.Low = valStr
		case fieldNameLast:
			tkr.Last = valStr
		case fieldNameBuy:
			tkr.Buy = valStr
		case fieldNameSell:
			tkr.Sell = valStr
		default:
			if !strings.HasPrefix(k, "vol_") {
				continue
			}

			volName := strings.Split(k, "_")
			if len(volName) != 2 {
				continue
			}
			if tkr.volumes == nil {
				tkr.volumes = make(map[string]string)
			}

			tkr.volumes[volName[1]] = valStr
		}
	}

	return nil
}

// func (tkr *Ticker) ToJson() (err error) {

// }
