package types

import (
	"encoding/json"
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

func (r Rate) Symbol() string {
	return r.From + "/" + r.To
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

type RateInfo struct {
	Price   string
	Updated time.Time
}

func NewRateInfo(price string) RateInfo {
	return RateInfo{
		Price:   price,
		Updated: time.Now(),
	}
}

func (ri RateInfo) Marshal() []byte {
	data, _ := json.Marshal(ri)
	return data
}

func (ri *RateInfo) Unmarshal(data []byte) error {
	return json.Unmarshal(data, ri)
}

type MarketRates map[Rate]RateInfo
