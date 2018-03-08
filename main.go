package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tjandrayana/cache-using-badger/module"
)

func main() {

	// Init the simple module
	m := module.Init()
	defer m.DB.Close()
	defer m.Txn.Discard()

	k := "try"
	v := "this"
	sec := 10 // set as 10 second

	// function to insert key value
	m.InsertKV(k, v, sec)

	var value string
	var err error

	// Get The value
	value, err = m.GetValue(k)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("value : ", value)

	// sleep used to test TTL of this key value
	time.Sleep(5 * time.Second)
	fmt.Println(" === after 5 second ==== ")

	value, err = m.GetValue(k)
	if err != nil {
		log.Println(err)
	}

	time.Sleep(6 * time.Second)

	fmt.Println(" === after 6 second ==== ")

	value, err = m.GetValue(k)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("==== Finish ==== ")

}
