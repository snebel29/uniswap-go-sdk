package uniswap

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

// Token represents an ERC20 token with a unique address and some metadata.
// TODO: To enforce read only, should Token be an interface instead?
type Token struct {
	*Currency
	ChainID ChainID
	Address common.Address
}

// NewToken return a new Token.
func NewToken(chainID ChainID, address string, decimals uint8, symbol, name string) (*Token, error) {
	if !common.IsHexAddress(address) {
		return nil, fmt.Errorf("address %q is invalid", address)
	}
	return &Token{
		Currency: NewCurrency(decimals, symbol, name),
		ChainID: chainID,
		Address: common.HexToAddress(address),
	}, nil
}

// Equals return true if the two tokens are equivalent, i.e. have the same chainId and address.
func (t *Token) Equals(token *Token) bool {
	if t == token {
		return true
	}
	return t.ChainID == token.ChainID && t.Address == token.Address
}

// SortsBefore return true if the address of t token sorts before the address of token.
func (t *Token) SortsBefore(token *Token) bool {
	// TODO: Do we really need this function method?
	return false
}

// TODO: Do we need to create currencyEquals function?
