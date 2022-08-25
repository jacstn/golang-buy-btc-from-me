package ext

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
)

func GetBTCPrice() string {
	client := binance.NewClient("", "")

	klines, err := client.NewKlinesService().Symbol("BTCUSDT").
		Interval("15m").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return "0"
	}

	return klines[len(klines)-1].Close
}
