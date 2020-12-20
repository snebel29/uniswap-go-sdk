package uniswap

// Token represents an ERC20 token with a unique address and some metadata.
type Token struct {
	Currency
	ChainID ChainID
	Address string
}

// NewToken return a new Token.
func NewToken(chainID ChainID, address string, decimals int, symbol, name string) Token {
	return Token{
		Currency: Currency{
			Decimals: decimals,
			Symbol: symbol,
			Name: name,
		},
		ChainID: chainID,
		Address: address,
	}
}
