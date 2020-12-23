package uniswap

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestToken_NewToken(t *testing.T) {
	for idx, f := range []struct {
		shouldErr bool
		chainID   ChainID
		address   string
		decimals  uint8
		symbol    string
		name      string
	}{
		// Start with a valid case.
		{false, MAINNET, "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599", 8, "WBTC", "Wrapped BTC"},
		// Invalid addresses.
		{true, MAINNET, "0xInvalidAddress", 8, "WBTC", "Wrapped BTC"},
		{true, MAINNET, "0x2260fac5e5542a773aa44fbcfedf7c193bc2c5999", 8, "WBTC", "Wrapped BTC"},
		{true, MAINNET, "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599z", 8, "WBTC", "Wrapped BTC"},
	} {
		msg := fmt.Sprintf("Test case at index %d failed", idx)
		token, err := NewToken(f.chainID, f.address, f.decimals, f.symbol, f.name)
		if f.shouldErr {
			assert.Error(t, err, msg)
			continue
		}
		require.NoError(t, err, msg)
		assert.Equal(t, token.ChainID, f.chainID, msg)
		assert.Equal(t, token.Address, common.HexToAddress(f.address), msg)
		assert.Equal(t, token.Decimals, f.decimals, msg)
		assert.Equal(t, token.Symbol, f.symbol, msg)
		assert.Equal(t, token.Name, f.name, msg)
	}
}

func TestToken_Equals(t *testing.T) {
	for idx, f := range []struct {
		shouldEqual   bool
		token0chainID ChainID
		token0address string
		token1chainID ChainID
		token1address string
	}{
		// Start with a valid case.
		{true, MAINNET, "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599", MAINNET, "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599"},
		// Different chain IDs.
		{false, MAINNET, "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599", ROPSTEN, "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599"},
		// Different addresses.
		{false, MAINNET, "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599", MAINNET, "0x6b175474e89094c44da98b954eedeac495271d0f"},
	} {
		msg := fmt.Sprintf("Test case at index %d failed", idx)
		token0, err := NewToken(f.token0chainID, f.token0address, 18, "", "")
		require.NoError(t, err, msg)
		token1, err := NewToken(f.token1chainID, f.token1address, 18, "", "")
		require.NoError(t, err, msg)
		require.True(t, token0.Equals(token1) == token1.Equals(token0), msg)
		assert.Equal(t, f.shouldEqual, token0.Equals(token1))
	}
}
