package uniswap

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	tokenDataCache = make(map[ChainID]map[common.Address]*tokenData)
)

type tokenData struct {
	decimals uint8
	symbol   string
	name     string
}

// FetchTokenData fetches information for a given token on the given chain, using the given ethers client.
func FetchTokenData(client *ethclient.Client, chainID ChainID, address string) (*Token, error) {
	var tokenData *tokenAddress0
	if td, ok := tokenDataCache[chainID][common.HexToAddress(address)]; ok {
		tokenData = td
	} else {
		// Fetch from the block chain.
		token, err := erc20.NewToken(c.tokenAddress0, client)
		if err != nil {
			return "", fmt.Errorf("failed to instantiate token0: %v", err)
		}
		tokenDecimals, err := token0.Decimals(nil)
		if err != nil {
			return "", fmt.Errorf("failed to get token0 decimals: %v", err)
		}
	}

	token, err := NewToken(chainID, address, decimals, symbol, name)
	if err != nil {
		return nil, fmt.Errorf("failed to create token: %v", err)
	}
	return token, nil
}

// FetchPairData fetches information about a pair and constructs a pair from the given two tokens.
func FetchPairData(client *ethclient.Client,) {
	// TODO: Implement
}
