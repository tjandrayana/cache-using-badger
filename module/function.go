package module

import (
	"log"
	"time"

	"github.com/dgraph-io/badger"
)

// InserKV means store key and value like store redis.io
// This function need key, value, and duration param
// If you dont want to set Expired just give the duration -1. So this mean unlimited store
func (m *Module) InsertKV(key, value string, duration int) error {
	var err error

	if duration == -1 {
		err = m.Txn.Set([]byte(key), []byte(value))
		if err != nil {
			log.Println(err)
		}
	} else {
		err = m.Txn.SetWithTTL([]byte(key), []byte(value), time.Duration(duration)*time.Second)
		if err != nil {
			log.Println(err)
		}
	}

	// Commit the transaction and check for error.
	err = m.Txn.Commit(nil)
	if err != nil {
		log.Println(err)
	}

	return err
}

// Get Value  will return the value of a key that stored in local disk data
// If the key not found and then will return empty string and error
func (m *Module) GetValue(key string) (string, error) {
	var value string
	err := m.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		val, err := item.Value()
		if err != nil {
			return err
		}
		value = string(val)
		return nil
	})
	if err != nil {
		log.Println(err)
	}

	return value, err
}
