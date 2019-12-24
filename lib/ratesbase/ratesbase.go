package ratesbase

import (
	"fmt"

	. "github.com/Croohand/crypto-monitor/lib/markets"
	. "github.com/Croohand/crypto-monitor/lib/types"

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

func Get(market Market, rate Rate) (RateInfo, error) {
	var rateInfo RateInfo
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(market))
		if b == nil {
			return fmt.Errorf("Can't find %s market in ratesbase", market)
		}
		return rateInfo.Unmarshal(b.Get([]byte(rate.Symbol())))
	})
	return rateInfo, err
}

func Set(market Market, rate Rate, rateInfo RateInfo) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(market))
		if err != nil {
			return err
		}
		return b.Put([]byte(rate.Symbol()), rateInfo.Marshal())
	})
}
