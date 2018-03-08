# cache-using-badger
This repository is used to try another way to cache the data besides redis. 
We can store the data by key, value with this optional. This is an optional way to cache data except redis. 
We can store using key, value and then set the TTL of our data. 
Unlike redis that use RAM, badger use our disk. 

The detail of badger can access here : 
https://github.com/dgraph-io/badger


# Init Module
m := module.Init()

# Store Key Value With Expired
    k := "try"
	v := "this"
	sec := 10 // set as 10 second expired

	var value string
	var err error

	// function to insert key value
	err = m.InsertKV(k, v, sec)
	if err != nil {
		log.Println(err)
	}

# Store Key Value Without Expired
    k := "try"
	v := "this"
	sec := -1 // -1 mean we not set the expired

	var value string
	var err error

	// function to insert key value
	err = m.InsertKV(k, v, sec)
	if err != nil {
		log.Println(err)
	}


# Get Value By Key
	value, err = m.GetValue(k)
	if err != nil {
		log.Println(err)
	}

