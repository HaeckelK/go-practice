// Quote Tracker (market symbols etc)
// A program which can go out and check the current value of stocks
// for a list of symbols entered by the user. The user can set how
// often the stocks are checked. For CLI, show whether the stock has
// moved up or down. Optional: If GUI, the program can show green up
// and red down arrows to show which direction the stock value has moved.

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

type Quote struct {
	code     string
	value    int
	movement int
}

type QuoteStore struct {
	quotes map[string][]*Quote
}

func (s *QuoteStore) AddQuote(q *Quote) {
	quotes := s.quotes[q.code]
	// If prior quote for this code exists then use value to calculate movement
	if len(quotes) > 0 {
		prior := quotes[len(quotes)-1]
		q.movement = q.value - prior.value
	}
	s.quotes[q.code] = append(s.quotes[q.code], q)
}

func (s *QuoteStore) AddQuotes(quotes []*Quote) {
	for _, quote := range quotes {
		s.AddQuote(quote)
		// TODO remove print
		fmt.Println(*quote)
	}
}

func NewQuoteStore() *QuoteStore {
	store := new(QuoteStore)
	store.quotes = make(map[string][]*Quote)
	return store
}

type Fetcher interface {
	Fetch(code string) (quote Quote, err error)
}

type TimeDelayFetcher struct {
	delay int
}

func (q TimeDelayFetcher) Fetch(code string) (quote Quote, err error) {
	time.Sleep(time.Duration(q.delay) * time.Second)
	quote = Quote{code: code, value: rand.Intn(100)}
	return
}

func getTargetCodes() []string {
	var codes []string
	for _, raw := range os.Args[1:] {
		codes = append(codes, strings.ToUpper(raw))
	}
	return codes
}

// TODO this was only here to eyeball the concurrency was working
func getQuote(fetcher Fetcher, code string) (quote Quote, err error) {
	fmt.Println(code, "start")
	quote, err = fetcher.Fetch(code)
	fmt.Println(code, "end")
	return
}

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
	// TODO frequency
	// TODO colour in terminal
	fmt.Println("Stock Quote Fetcher")
	fetcher := TimeDelayFetcher{delay: 0}
	codes := getTargetCodes()
	store := NewQuoteStore()

	fmt.Println("Codes Provided:")
	fmt.Println(codes)

	fmt.Println("Fetching Quotes")
	quotes := getQuotes(fetcher, codes)
	store.AddQuotes(quotes)

	quotes = getQuotes(fetcher, codes)
	store.AddQuotes(quotes)

}
