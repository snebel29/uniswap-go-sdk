package uniswap

// ETHER represents ether.
var ETHER = Currency{Decimals: 18, Symbol: "ETH", Name: "Ether"}

// Currency is any fungible financial instrument on Ethereum, including Ether and all ERC20 tokens.
type Currency struct {
	// TODO: use bigint?
	Decimals int
	Symbol   string
	Name     string
}

// NewCurrency return a new currency.
func NewCurrency(decimals int, symbol, name string) Currency {
	// TODO: Validate decimals is within solidity types uint8 and uint256.
	return Currency{
		Decimals: decimals,
		Symbol:   symbol,
		Name:     name,
	}
}
