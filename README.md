# Cleanaero
A small program to help clean aerospike server databases during testing. It is not meant for production or maintanance, but simply to remove old and unnecessary testing data that may be left over.

## Usage
```
-host string
      Database server hostname. (default "localhost")
-ns string
      Namespace to access within the database. (default "test")
-port int
      Database server port number. (default 3000)
-sets string
      Sets to wipe clean in the database.
-v    Make program output verbose.
```

## Dependencies
Requires the [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) client library.

### Installing aerospike-client-go
Get the client in your `GOPATH`
```
go get github.com/aerospike/aerospike-client-go
```

## Important
The program relies on being able to query the database and receiving primary keys. Currently there is no other method to remove values from aerospike sets, and the server will only return primary keys if they were specifically sent to the server when the data was added. The exact implementation between different versions of the client library can vary. However, in the Go version this is done through the use of the [WritePolicy](https://godoc.org/github.com/aerospike/aerospike-client-go#WritePolicy) datatype, by setting `SendKey` to `True`.

### Example
```Go
package main

import (
	db "github.com/aerospike/aerospike-client-go"
)

func test() {
	policy := db.NewWritePolicy(0, 0)
	policy.SendKey = true

	key, _ := db.NewKey("test", "myset", "mykey")
	bName := db.NewBin("Name", "Jordan")

	err := DBConn.PutBins(policy, key, bName)
	panicOnError(err)
}
```

## TODO List
- [ ] Add support for UserPolicies.
- [ ] Add test data support.
