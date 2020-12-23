package uniswap

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/snebel29/uniswap-go-sdk/contracts/erc20"
)

type tokenDataCache map[ChainID]map[common.Address]*tokenData

func (tdc tokenDataCache) addTokenData(chainID ChainID, addr common.Address, td *tokenData) {
	if tdc[chainID] == nil {
		tdc[chainID] = make(map[common.Address]*tokenData)
	}
	tdc[chainID][addr] = td
}

type tokenData struct {
	decimals uint8
	symbol   string
	name     string
}

// Fetcher represents a fetcher.
type Fetcher struct {
	// TODO: Alias client, common, bind, etc. for users to not have to import other libraries.
	client         *ethclient.Client
	newErc20Caller func(common.Address, bind.ContractBackend) (erc20.ReadOnlyContract, error)
	tokenDataCache tokenDataCache
}

// NewFetcher creates a new fetcher.
func NewFetcher(client *ethclient.Client) *Fetcher {
	return &Fetcher{
		// TODO: Does a nil client cause dowsantream issues?
		client: client,
		newErc20Caller: func(addr common.Address, client bind.ContractBackend) (erc20.ReadOnlyContract, error) {
			return erc20.NewErc20Caller(addr, client)
		},
		tokenDataCache: make(tokenDataCache),
	}
}

// FetchTokenData fetches information for a given token on the given chain, using the given ethers client.
func (f *Fetcher) FetchTokenData(chainID ChainID, address string) (*Token, error) {
	var td *tokenData
	addr := common.HexToAddress(address)
	if cachedTD, ok := f.tokenDataCache[chainID][addr]; ok {
		td = cachedTD
	} else {
		// Fetch from the block chain.
		token, err := f.newErc20Caller(addr, f.client)
		if err != nil {
			return nil, fmt.Errorf("failed to bound to erc20 token contract: %v", err)
		}
		// TODO: Should we set some default CallOpts?
		var callOpts *bind.CallOpts
		decimals, err := token.Decimals(callOpts)
		if err != nil {
			return nil, fmt.Errorf("failed to get token decimals: %v", err)
		}
		symbol, err := token.Symbol(callOpts)
		if err != nil {
			return nil, fmt.Errorf("failed to get token symbol: %v", err)
		}
		name, err := token.Name(callOpts)
		if err != nil {
			return nil, fmt.Errorf("failed to get token name: %v", err)
		}
		td = &tokenData{
			decimals: decimals,
			symbol:   symbol,
			name:     name,
		}
	}

	token, err := NewToken(chainID, address, td.decimals, td.symbol, td.name)
	if err != nil {
		return nil, fmt.Errorf("failed to create token: %v", err)
	}
	f.tokenDataCache.addTokenData(chainID, addr, td)
	return token, nil
}

// FetchPairData fetches information about a pair and constructs a pair from the given two tokens.
func FetchPairData(client *ethclient.Client) {
	// TODO: Implement
}
