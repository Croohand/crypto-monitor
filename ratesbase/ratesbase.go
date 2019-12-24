package ratesbase

import (
	"fmt"

	"github.com/Croohand/crypto-monitor/markets"
	"github.com/Croohand/crypto-monitor/types"

	bolt "go.etcd.io/bbolt"
)

var db *bolt.DB

func Open(path string) error {
	var err error
	db, err = bolt.Open(path, 0600, nil)
	return err
}

func Close() error {
	return db.Close()
}

func Get(market markets.Market, rate types.Rate) (*types.RateInfo, error) {
	rateInfo := new(types.RateInfo)
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(market))
		if b == nil {
			return fmt.Errorf("Can't find %s market in ratesbase", market)
		}
		return rateInfo.Unmarshal(b.Get([]byte(rate.Symbol())))
	})
	return rateInfo, err
}

func Set(market markets.Market, rateInfo *types.RateInfo) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(market))
		if err != nil {
			return err
		}
		return b.Put([]byte(rateInfo.Symbol()), rateInfo.Marshal())
	})
}
