package main

import (
	"strconv"
	"testing"
)

// TODO split into separate tests
func TestQuoteStoreAddQuote(t *testing.T) {
	// Given a new QuoteStore
	store := NewQuoteStore()
	// When adding multiple quotes
	q1 := Quote{code: "code", value: 999}
	q2 := Quote{code: "code", value: 1000}
	store.AddQuote(&q1)
	store.AddQuote(&q2)
	// Then length of store.quotes is number of quotes added
	if len(store.quotes["code"]) != 2 {
		t.Errorf("Expected len 2, got %s", strconv.Itoa(len(store.quotes["code"])))
	}
	// Then order of adding is preserved
	if store.quotes["code"][1] != &q2 {
		t.Error("Final quote does not agree to latest quote added.")
	}
	// Then q2 movement is calculated
	if q2.movement != (q2.value - q1.value) {
		// TODO proper error message
		t.Errorf("Movement does not equal q2 - q1 value")
	}
}
