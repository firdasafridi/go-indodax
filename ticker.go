package indodax

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

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

	high, _ := strconv.ParseFloat(tkr.Ticker.High, 64)

	low, _ := strconv.ParseFloat(tkr.Ticker.Low, 64)

	var assetVolume, baseVolume float64 = 0, 0
	volName := strings.Split(pairName, "_")
	if len(tkr.Ticker.volumes) < 2 {
		return ticker, err
	}
	if v, ok := tkr.Ticker.volumes[volName[0]]; ok {
		assetVolume, _ = strconv.ParseFloat(v, 64)
	}
	if v, ok := tkr.Ticker.volumes[volName[1]]; ok {
		baseVolume, _ = strconv.ParseFloat(v, 64)
	}

	last, _ := strconv.ParseFloat(tkr.Ticker.Last, 64)

	buy, _ := strconv.ParseFloat(tkr.Ticker.Buy, 64)

	sell, _ := strconv.ParseFloat(tkr.Ticker.Sell, 64)

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
		case fieldNameAmount:
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
