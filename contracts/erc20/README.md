# erc20
This package implement the erc20  interface.

## Re-generate this package

### Install abigen
Assuming that the module version of `go-ethereum` in the SDK is `v1.9.25`.
```go
$ go get -u github.com/ethereum/go-ethereum
$ cd $GOPATH/src/github.com/ethereum/go-ethereum
$ git checkout v1.9.25
$ make 
$ make devtools
```

Check that the installation went well.
```
$ abigen --version
abigen version 1.9.25-stable
```

### Create go bindings from abi file
```shell
$ abigen --abi=erc20.abi --pkg=erc20 --out=erc20.go
```
