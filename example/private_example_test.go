package indodax

import (
	"fmt"
	"log"
	"time"
	
	"github.com/firdasafridi/go-indodax"
)

//
// Example trade buy
//
func ExampleClient_TradeBuy() {
	cl, err := indodax.NewClient(
		"Key",
		"Secret",
	)
	if err != nil {
		log.Println(err)
		return
	}

	var pairName string = "usdt_idr"
	var price float64 = 12000
	var amount float64 = 50000

	tradeBuy, err := cl.TradeBuy(pairName, price, amount)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("TradeBuy response: %+v\n", tradeBuy)
}

func ExampleClient_TradeSell() {
	cl, err := indodax.NewClient(
		"Key",
		"Secret",
	)
	if err != nil {
		log.Println(err)
		return
	}

	var pairName string = "usdt_idr"
	var price float64 = 12000
	var amount float64 = 50000

	tradeSell, err := cl.TradeSell(pairName, price, amount)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("TradeSell response: %+v\n", tradeSell)
}

func ExampleClient_CancelOrderBuy() {
	cl, err := indodax.NewClient(
		"Key",
		"Secret",
	)
	if err != nil {
		log.Println(err)
		return
	}

	var pairName string = "usdt_idr"
	var orderId int64 = 12345

	cancelBuy, err := cl.CancelOrderBuy(pairName, orderId)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("CancelOrderBuy response: %+v\n", cancelBuy)
}

func ExampleClient_CancelOrderSell() {
	cl, err := indodax.NewClient(
		"Key",
		"Secret",
	)
	if err != nil {
		log.Println(err)
		return
	}

	var pairName string = "usdt_idr"
	var orderId int64 = 12345

	cancelSell, err := cl.CancelOrderSell(pairName, orderId)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("CancelOrderSell response: %+v\n", cancelSell)
}

func ExampleClient_GetOrder() {
	cl, err := indodax.NewClient(
		"Key",
		"Secret",
	)
	if err != nil {
		log.Println(err)
		return
	}

	var pairName string = "usdt_idr"
	var orderId int64 = 12345

	getOrder, err := cl.GetOrder(pairName, orderId)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("GetOrder response: %+v\n", getOrder)
}

func ExampleClient_TradeHitory() {
	cl, err := indodax.NewClient(
		"Key",
		"Secret",
	)
	if err != nil {
		log.Println(err)
		return
	}

	// pair name is required
	var pairName string = "usdt_idr"

	// set time is optional
	sinceTime := time.Date(
	    2000, 11, 17, 20, 34, 58, 651387237, time.UTC)
	endTime := time.Date(
		2020, 11, 17, 20, 34, 58, 651387237, time.UTC)
	
	// count, start trade id, and end trade id is optional 
	var count, startTradeId, endTradeId int64 = 10, 20, 30

	// order is optional
	var sortOrder string = "asc"
	tradeHitory, err := cl.TradeHitory(pairName, count, startTradeId, endTradeId, sortOrder, &sinceTime, &endTime)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("TradeHitory response: %+v\n", tradeHitory)
}

func ExampleClient_OrderHistory() {
	cl, err := indodax.NewClient(
		"Key",
		"Secret",
	)
	if err != nil {
		log.Println(err)
		return
	}

	var pairName string = "usdt_idr"
	var count, from int64 = 10, 20

	orderHistory, err := cl.OrderHistory(pairName, count, from)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("OrderHistory response: %+v\n", orderHistory)
}

func ExampleClient_AllOpenOrders() {
	cl, err := indodax.NewClient(
		"Key",
		"Secret",
	)
	if err != nil {
		log.Println(err)
		return
	}

	allOrdes, err := cl.AllOpenOrders()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("AllOpenOrders response: %+v\n", allOrdes)
}

func ExampleClient_TransHistory() {
	cl, err := indodax.NewClient(
		"Key",
		"Secret",
	)
	if err != nil {
		log.Println(err)
		return
	}

	transHistory, err := cl.TransHistory()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("TransHistory response: %+v\n", transHistory)
}

func ExampleClient_GetInfo() {
	cl, err := indodax.NewClient(
		"Key",
		"Secret",
	)
	if err != nil {
		log.Println(err)
		return
	}

	getInfo, err := cl.GetInfo()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("GetInfo response: %+v\n", getInfo)
}
