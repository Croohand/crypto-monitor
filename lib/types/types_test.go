package types

import (
	"testing"
)

func TestParsingBad(t *testing.T) {
	badSymbols := []string{
		"A/b",
		"AB",
		"A//B",
		"/B",
		"A/",
	}

	for _, symbol := range badSymbols {
		if CanParse(symbol) {
			t.Errorf("Can parse symbol %s while shouldn't", symbol)
		}
	}

	if len(ParseRates(badSymbols)) > 0 {
		t.Error("Non-empty result of ParseRates on bad symbols")
	}
}

func TestParsingGood(t *testing.T) {
	goodSymbols := []string{
		"A/B",
		"A/BC",
		"AB/C",
	}

	for _, symbol := range goodSymbols {
		if !CanParse(symbol) {
			t.Errorf("Can't parse symbol %s while should", symbol)
		}
	}

	goodRates := []Rate{
		Rate{"A", "B"},
		Rate{From: "A", To: "BC"},
		Rate{From: "AB", To: "C"},
	}

	if !testRatesEq(ParseRates(goodSymbols), goodRates) {
		t.Errorf("ParseRates(%#v) != %#v", goodSymbols, goodRates)
	}
}

func testRatesEq(a, b []Rate) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
