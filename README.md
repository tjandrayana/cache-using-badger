# cache-using-badger
This repository is used to try another way to cache the data besides redis. 
We can store the data by key, value with this optional. This is an optional way to cache data except redis. 
We can store using key, value and then set the TTL of our data. 
Unlike redis that use RAM, badger use our disk. 

The detail of badger can access here : https://github.com/dgraph-io/badger
