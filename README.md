# Cleanaero
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

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
Requires the [aerospike client go](https://github.com/aerospike/aerospike-client-go) client library.

### Installing aerospike-client-go
Get the client in your `GOPATH`
```
go get github.com/aerospike/aerospike-client-go
```

## TODO List
- [ ] Add support for UserPolicies.
- [ ] Add test data support.
