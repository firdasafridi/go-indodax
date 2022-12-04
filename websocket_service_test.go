package indodax

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"sync"
	"sync/atomic"
	"testing"
)

func TestWsOrderBookServe(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	aEvent := atomic.Pointer[OrderBookEvent]{}
	aErr := atomic.Pointer[error]{}

	eventHandler := func(event *OrderBookEvent) {
		aEvent.Store(event)
		wg.Done()
	}
	errHandler := func(err error) {
		aErr.Store(&err)
		wg.Done()
	}

	doneC, stopC, err := WsOrderBookServe(OBS("btc", "idr"), eventHandler, errHandler)
	assert.NoError(t, err)

	wg.Wait()

	stopC <- struct{}{}
	<-doneC

	if pErr := aErr.Load(); pErr != nil {
		require.NoError(t, *pErr)
	}

	e := aEvent.Load()
	assert.NotNil(t, e)
	assert.Equal(t, "btcidr", e.Pair, "pair")
	assert.NotEmpty(t, e.Ask, "ask")
	assert.NotEmpty(t, e.Bid, "bid")
}

func OBS(base, quote string) OrderBookSymbol {
	return OrderBookSymbol{
		Base:  base,
		Quote: quote,
	}
}

func Test_toOrderBookEvent(t *testing.T) {
	c, err := os.ReadFile("orderbook_response.json")
	require.NoError(t, err)

	e, offset, err, isTargetedMessage := toOrderBookEvent("market:order-book-btcidr", OBS("btc", "idr"), 0, c)
	require.NoError(t, err)
	require.True(t, isTargetedMessage)
	require.NotNil(t, e)
	assert.Equal(t, 6563500, int(offset))

	if assert.Len(t, e.Ask, 2) {
		assertOrderBookEntry(t, e.Ask[0], "0.74041081", "195433655", "263953000")
		assertOrderBookEntry(t, e.Ask[1], "0.00762000", "2011337", "263955000")
	}

	if assert.Len(t, e.Bid, 2) {
		assertOrderBookEntry(t, e.Bid[0], "0.01769092", "4669554", "263952000")
		assertOrderBookEntry(t, e.Bid[1], "0.00037693", "99490", "263949000")
	}
}

func assertOrderBookEntry(t *testing.T, got OrderBookEntry, wantBase, wantQuote, wantPrice string) {
	assert.Equal(t, wantBase, fmt.Sprintf("%0.8f", got.BaseVolume), "base")
	assert.Equal(t, wantQuote, fmt.Sprintf("%0.0f", got.QuoteVolume), "quote")
	assert.Equal(t, wantPrice, fmt.Sprintf("%0.0f", got.Price), "price")
}

func Test_firstError(t *testing.T) {
	err1 := errors.New("hello")
	err2 := errors.New("world")

	type args struct {
		fst error
		snd error
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{"nil-nil", args{nil, nil}, nil},
		{"nil-err", args{nil, err2}, err2},
		{"err-nil", args{err1, nil}, err1},
		{"err-err", args{err1, err2}, err1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := firstError(tt.args.fst, tt.args.snd)

			assert.Equal(t, tt.wantErr, got, fmt.Sprintf("firstError(%v, %v)", tt.args.fst, tt.args.snd))
		})
	}
}
