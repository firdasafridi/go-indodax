# GO-Indodax - A Library trading platform Indodax using Go Language
- [Description](#description)
- [Features](#features)
- [Documentation](#documentation)
- [godoc](https://godoc.org/github.com/firdasafridi/go-indodax)
- [Example](#example)

## Description
This is unofficial library golang for Indodax trading platform.

## Features

Public Endpoint
- Ticker
- Depth (Order Book)
- Trades (Trade History)

Private Endpoint
- Get Information User
- Withdraw/Deposit History
- Trading History
- Order History
- Open Orders
- Trade
- Withdraw (Comming Soon)

## Example

Public Endpoint 
``` go
func main() {
    cl, err := indodax.NewClient(
		"",
		"",
	)
	ticker, err := cl.GetTicker("ten_idr")
	if err != nil {
		fmt.Println(err)
	}
    fmt.Println(ticker)
}
```

Private Endpoint 
```go
func main() {
    cl, err := indodax.NewClient(
		"key", 
		"secret", 
	)
	tradeBuy, err := cl.TradeBuy("usdt_idr", 12000, 50000)
	if err != nil {
		fmt.Println(err)
	}
	jsonToString(tradeBuy)
}
```
