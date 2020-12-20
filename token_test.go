package uniswap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	chainID := MAINNET
	address := "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599"
	decimals := 8
	symbol := "WBTC"
	name := "Wrapped BTC"
	token := NewToken(chainID, address, decimals, symbol, name)
	assert.Equal(t, token.ChainID, chainID)
	assert.Equal(t, token.Address, address)
	assert.Equal(t, token.Decimals, decimals)
	assert.Equal(t, token.Symbol, symbol)
	assert.Equal(t, token.Name, name)
}
