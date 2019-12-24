package types

import (
	"regexp"
	"time"
)

var rateSymbolRegexp *regexp.Regexp

func init() {
	rateSymbolRegexp = regexp.MustCompile("([A-Z]+)/([A-Z]+)")
}

type Rate struct {
	From, To string
}

type RateInfo struct {
	Rate
	Price   string
	Updated time.Time
}

func NewRateInfo(rate Rate, price string) RateInfo {
	return RateInfo{
		Rate:    rate,
		Price:   price,
		Updated: time.Now(),
	}
}

func CanParse(symbol string) bool {
	return rateSymbolRegexp.MatchString(symbol)
}

func ParseRates(symbols []string) []Rate {
	rates := []Rate{}
	for _, symbol := range symbols {
		if CanParse(symbol) {
			match := rateSymbolRegexp.FindStringSubmatch(symbol)
			rates = append(rates, Rate{match[1], match[2]})
		}
	}
	return rates
}
