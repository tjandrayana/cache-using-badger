package module

import (
	"log"

	"github.com/dgraph-io/badger"
)

type Module struct {
	DB  *badger.DB
	Txn *badger.Txn
}

const dir string = "db/badger"

func Init() *Module {
	opts := badger.DefaultOptions
	opts.Dir = dir
	opts.ValueDir = dir
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}

	return &Module{
		DB:  db,
		Txn: db.NewTransaction(true),
	}
}
