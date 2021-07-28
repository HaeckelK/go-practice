// Quote Tracker (market symbols etc)
// A program which can go out and check the current value of stocks
// for a list of symbols entered by the user. The user can set how
// often the stocks are checked. For CLI, show whether the stock has
// moved up or down. Optional: If GUI, the program can show green up
// and red down arrows to show which direction the stock value has moved.

package main

import (
	"fmt"
	"sync"
	"time"
)

type Quote struct {
	code  string
	value int
}

type Fetcher interface {
	Fetch(code string) (quote Quote, err error)
}

type TimeDelayFetcher struct {
	delay int
}

func (q TimeDelayFetcher) Fetch(code string) (quote Quote, err error) {
	time.Sleep(time.Duration(q.delay) * time.Second)
	return
}

func getTargetCodes() []string {
	return []string{"AMC", "AAPL"}
}

func getQuote(fetcher Fetcher, code string) (quote Quote, err error) {
	fmt.Println(code, "start")
	quote, err = fetcher.Fetch(code)
	fmt.Println(code, "end")
	return
}

// No concurrency
// func getQuotes(fetcher Fetcher, codes []string) (quotes []*Quote) {
// 	for _, code := range codes {
// 		quote, _ := getQuote(fetcher, code)
// 		quotes = append(quotes, &quote)
// 	}
// 	return
// }

func getQuotes(fetcher Fetcher, codes []string) (quotes []*Quote) {
	var wg sync.WaitGroup
	for _, code := range codes {
		wg.Add(1)
		go func(code string) {
			quote, _ := getQuote(fetcher, code)
			quotes = append(quotes, &quote)
			wg.Done()
		}(code)
	}
	wg.Wait()
	return
}

func main() {
	// TODO UI for quote codes
	// TODO frequency
	// TODO colour in terminal
	fmt.Println("Stock Quote Fetcher")
	fetcher := TimeDelayFetcher{delay: 2}
	codes := getTargetCodes()
	fmt.Println("Codes Provided:")
	fmt.Println(codes)

	fmt.Println("Fetching Quotes")
	quotes := getQuotes(fetcher, codes)
	fmt.Println(quotes)
	time.Sleep(60 * time.Second)

}
