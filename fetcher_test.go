package uniswap

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/snebel29/uniswap-go-sdk/contracts/erc20"
)

// TODO: Create some live test scenario in some of the test network.

func TestFetcher_FetchTokenData(t *testing.T) {
	decimals := uint8(18)
	symbol := "WBTC"
	name := "Wrapped BTC"
	address := "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599"
	// We can pass a nil client because erc20 caller contract wil be mocked.
	fetcher := NewFetcher(nil)

	newErc20CallerThatReturnError := func(addr common.Address, client bind.ContractBackend) (erc20.ReadOnlyContract, error) {
		return &erc20.ReadOnlyContractMock{}, fmt.Errorf("test error")
	}
	fetcher.newErc20Caller = newErc20CallerThatReturnError
	_, err := fetcher.FetchTokenData(MAINNET, address)
	assert.Error(t, err, "When erc20 caller return error FetchTokenData should do as well")

	decimalsThatReturnError := func(addr common.Address, client bind.ContractBackend) (erc20.ReadOnlyContract, error) {
		mock := &erc20.ReadOnlyContractMock{}
		mock.OnDecimals(func(*bind.CallOpts) (uint8, error) { return uint8(0), fmt.Errorf("test error") })
		return mock, nil
	}
	fetcher.newErc20Caller = decimalsThatReturnError
	_, err = fetcher.FetchTokenData(MAINNET, address)
	assert.Error(t, err, "When decimals return error FetchTokenData should do as well")

	symbolThatReturnError := func(addr common.Address, client bind.ContractBackend) (erc20.ReadOnlyContract, error) {
		mock := &erc20.ReadOnlyContractMock{}
		mock.OnSymbol(func(*bind.CallOpts) (string, error) { return "", fmt.Errorf("test error") })
		return mock, nil
	}
	fetcher.newErc20Caller = symbolThatReturnError
	_, err = fetcher.FetchTokenData(MAINNET, address)
	assert.Error(t, err, "When symbol return error FetchTokenData should do as well")

	nameThatReturnError := func(addr common.Address, client bind.ContractBackend) (erc20.ReadOnlyContract, error) {
		mock := &erc20.ReadOnlyContractMock{}
		mock.OnName(func(*bind.CallOpts) (string, error) { return "", fmt.Errorf("test error") })
		return mock, nil
	}
	fetcher.newErc20Caller = nameThatReturnError
	_, err = fetcher.FetchTokenData(MAINNET, address)
	assert.Error(t, err, "When name return error FetchTokenData should do as well")

	mock := &erc20.ReadOnlyContractMock{}
	successfulERC20Fetch := func(addr common.Address, client bind.ContractBackend) (erc20.ReadOnlyContract, error) {
		mock.OnDecimals(func(*bind.CallOpts) (uint8, error) { return decimals, nil })
		mock.OnSymbol(func(*bind.CallOpts) (string, error) { return symbol, nil })
		mock.OnName(func(*bind.CallOpts) (string, error) { return name, nil })
		return mock, nil
	}
	fetcher.newErc20Caller = successfulERC20Fetch
	token, err := fetcher.FetchTokenData(MAINNET, "badAddress")
	require.Error(t, err, "Bad address should produce an error")
	require.Equal(t, 0, len(fetcher.tokenDataCache), "At this point token data cache should still be empty")

	token, err = fetcher.FetchTokenData(MAINNET, address)
	require.NoError(t, err, "When no function return errors there should be no errors")
	assert.Equal(t, decimals, token.Decimals)
	assert.Equal(t, symbol, token.Symbol)
	assert.Equal(t, name, token.Name)
	assert.Equal(t, 1, len(fetcher.tokenDataCache), "By now the token data should have been cached")
	assert.EqualValues(t, 2, mock.DecimalsCalled(), "FetchTokenData was successfully called twice")
	assert.EqualValues(t, 2, mock.SymbolCalled(), "FetchTokenData was successfully called twice")
	assert.EqualValues(t, 2, mock.NameCalled(), "FetchTokenData was successfully called twice")

	token, err = fetcher.FetchTokenData(MAINNET, address)
	require.NoError(t, err, "When no function return errors there should be no errors")
	assert.EqualValues(t, 2, mock.DecimalsCalled(), "When token data is cached there should be no calls")
	assert.EqualValues(t, 2, mock.SymbolCalled(), "When token data is cached there should be no calls")
	assert.EqualValues(t, 2, mock.NameCalled(), "When token data is cached there should be no calls")
}
