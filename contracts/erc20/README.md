# erc20
This package implement the erc20  interface.


## Re-generate this package

### Install abigen
```go
$ go get -u github.com/ethereum/go-ethereum
$ cd $GOPATH/src/github.com/ethereum/go-ethereum
$ make 
$ make devtools
```
### Create go bindings from abi file
```shell
$ abigen --abi=erc20.abi --pkg=token --out=erc20.go
```
