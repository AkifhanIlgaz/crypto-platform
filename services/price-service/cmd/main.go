package main

import (
	"crypto-platform/shared/config"
	"fmt"
	"log"
	"os"
	"time"

	ccxt "github.com/ccxt/ccxt/go/v4"
)

func main() {
	_, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		os.Exit(1)
	}

	// println("start kuku")

	kucoin := ccxt.NewKucoin(nil)
	// binance := ccxt.NewBinance(nil)

	// kucoin.FetchOHLCV("XAGX/USDT")

	if _, err := kucoin.LoadMarkets(); err != nil {
		log.Fatalf("[%s] Piyasalar yüklenirken hata oluştu: %v", "Kucoin", err)
	}

	// if _, err := binance.LoadMarkets(); err != nil {
	// 	log.Fatalf("[%s] Piyasalar yüklenirken hata oluştu: %v", "Binanceß", err)
	// }

	symbol := "XAGX/USDT"

	ticker, err := kucoin.FetchTicker(symbol)
	if err != nil {
		log.Fatalf("[%s] Ticker verisi çekilemedi: %v", "Kucoin", err)
	}

	printTicker(ticker, "Kucoin")
	// if ticker.Last != nil {
	// 	fmt.Printf("[%s] Son Fiyat: %.2f USDT\n", "Kucoin", *ticker.Last)
	// }
	// if ticker.High != nil {
	// 	fmt.Printf("[%s] 24s En Yüksek: %.2f USDT\n", "Kucoin", *ticker.High)
	// }
	// if ticker.Low != nil {
	// 	fmt.Printf("[%s] 24s En Düşük: %.2f USDT\n", "Kucoin", *ticker.Low)
	// }
	// if ticker.BaseVolume != nil {
	// 	fmt.Printf("[%s] Hacim (BTC): %.4f\n", "Kucoin", *ticker.BaseVolume)
	// }

	// ohlcv, err := kucoin.FetchOHLCV("BTC/USDT", ccxt.WithFetchOHLCVTimeframe(`1d`), ccxt.WithFetchOHLCVLimit(2))
	// if err != nil {
	// 	log.Fatalf("[%s] OHLCV verisi çekilemedi: %v", "Kucoin", err)
	// }

	// lastCandle := ohlcv[0]
	// fmt.Printf("[%s] : %s\n", "Kucoin", time.UnixMilli(lastCandle.Timestamp))
	// fmt.Printf("[%s] Son Zamanda Açılış: %.2f USDT\n", "Kucoin", lastCandle.Open)
	// fmt.Printf("[%s] Son Zamanda Fiyat: %.2f USDT\n", "Kucoin", lastCandle.Close)
	// fmt.Printf("[%s] Son Zamanda Yüksek: %.2f USDT\n", "Kucoin", lastCandle.High)
	// fmt.Printf("[%s] Son Zamanda Düşük: %.2f USDT\n", "Kucoin", lastCandle.Low)
	// fmt.Printf("[%s] Son Zamanda Hacim (BTC): %.4f\n", "Kucoin", lastCandle.Volume)

}

func printTicker(ticker ccxt.Ticker, exchange string) {
	if ticker.Timestamp != nil {
		ts := *ticker.Timestamp
		fmt.Printf("[%s] Timestamp: %d (Unix ms, UTC: %s)\n", exchange, ts, time.UnixMilli(ts).UTC().Format(time.RFC3339))
	} else {
		fmt.Printf("[%s] Timestamp: n/a (Unix ms)\n", exchange)
	}

	fmt.Printf("[%s] Symbol: %s (Piyasa sembolü)\n", exchange, formatStringPtr(ticker.Symbol))
	fmt.Printf("[%s] Datetime: %s (İnsan okunur zaman)\n", exchange, formatStringPtr(ticker.Datetime))
	fmt.Printf("[%s] High: %s (Dönem en yüksek fiyatı)\n", exchange, formatFloatPtr(ticker.High))
	fmt.Printf("[%s] Low: %s (Dönem en düşük fiyatı)\n", exchange, formatFloatPtr(ticker.Low))
	fmt.Printf("[%s] Bid: %s (En iyi alış fiyatı)\n", exchange, formatFloatPtr(ticker.Bid))
	fmt.Printf("[%s] BidVolume: %s (En iyi alıştaki miktar)\n", exchange, formatFloatPtr(ticker.BidVolume))
	fmt.Printf("[%s] Ask: %s (En iyi satış fiyatı)\n", exchange, formatFloatPtr(ticker.Ask))
	fmt.Printf("[%s] AskVolume: %s (En iyi satıştaki miktar)\n", exchange, formatFloatPtr(ticker.AskVolume))
	fmt.Printf("[%s] Vwap: %s (Hacim ağırlıklı ortalama fiyat)\n", exchange, formatFloatPtr(ticker.Vwap))
	fmt.Printf("[%s] Open: %s (Dönem açılış fiyatı)\n", exchange, formatFloatPtr(ticker.Open))
	fmt.Printf("[%s] Close: %s (Dönem kapanış fiyatı)\n", exchange, formatFloatPtr(ticker.Close))
	fmt.Printf("[%s] Last: %s (Son işlem fiyatı)\n", exchange, formatFloatPtr(ticker.Last))
	fmt.Printf("[%s] PreviousClose: %s (Önceki kapanış)\n", exchange, formatFloatPtr(ticker.PreviousClose))
	fmt.Printf("[%s] Change: %s (Mutlak değişim)\n", exchange, formatFloatPtr(ticker.Change))
	fmt.Printf("[%s] Percentage: %s (Yüzde değişim)\n", exchange, formatFloatPtr(ticker.Percentage))
	fmt.Printf("[%s] Average: %s (Ortalama fiyat)\n", exchange, formatFloatPtr(ticker.Average))
	fmt.Printf("[%s] BaseVolume: %s (Baz para hacmi)\n", exchange, formatFloatPtr(ticker.BaseVolume))
	fmt.Printf("[%s] QuoteVolume: %s (Karşı para hacmi)\n", exchange, formatFloatPtr(ticker.QuoteVolume))

	if ticker.Info != nil {
		fmt.Printf("[%s] Info: %+v (Ham borsa cevabı)\n", exchange, ticker.Info)
	}
}
func formatStringPtr(p *string) string {
	if p == nil {
		return "n/a"
	}
	return *p
}

func formatFloatPtr(p *float64) string {
	if p == nil {
		return "n/a"
	}
	return fmt.Sprintf("%.8f", *p)
}

func formatInt64Ptr(p *int64) string {
	if p == nil {
		return "n/a"
	}
	return fmt.Sprintf("%d", *p)
}
