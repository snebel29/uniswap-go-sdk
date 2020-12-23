package erc20

import (
	"sync/atomic"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// ReadOnlyContract defines an ERC20 read-only contract.
type ReadOnlyContract interface {
	Decimals(*bind.CallOpts) (uint8, error)
	Symbol(*bind.CallOpts) (string, error)
	Name(*bind.CallOpts) (string, error)
}

// ReadOnlyContractMock is a read only contract mock.
type ReadOnlyContractMock struct {
	decimalsFn func(*bind.CallOpts) (uint8, error)
	decimalsCalled uint64

	symbolFn func(*bind.CallOpts) (string, error)
	symbolCalled uint64

	nameFn func(*bind.CallOpts) (string, error)
	nameCalled uint64
}

// OnDecimals replaces makes mock to call fn when Decimals function is called.
func (m *ReadOnlyContractMock) OnDecimals(fn func(*bind.CallOpts) (uint8, error)) {
	m.decimalsFn = fn
}

// DecimalsCalled return the number of times the function was called in this instance.
func (m *ReadOnlyContractMock) DecimalsCalled() uint64 {
	return atomic.LoadUint64(&m.decimalsCalled)
}

// Decimals implement the ReadOnlyContract interface.
func (m *ReadOnlyContractMock) Decimals(callOpts *bind.CallOpts) (uint8, error) {
	atomic.AddUint64(&m.decimalsCalled, 1)
	if m.decimalsFn != nil {
		return m.decimalsFn(callOpts)
	}
	return 0, nil
}

// OnSymbol replaces makes mock to call fn when Symbol function is called.
func (m *ReadOnlyContractMock) OnSymbol(fn func(*bind.CallOpts) (string, error)) {
	m.symbolFn = fn
}

// SymbolCalled return the number of times the function was called in this instance.
func (m *ReadOnlyContractMock) SymbolCalled() uint64 {
	return atomic.LoadUint64(&m.symbolCalled)
}

// Symbol implement the ReadOnlyContract interface.
func (m *ReadOnlyContractMock) Symbol(callOpts *bind.CallOpts) (string, error) {
	atomic.AddUint64(&m.symbolCalled, 1)
	if m.symbolFn != nil {
		return m.symbolFn(callOpts)
	}
	return "", nil
}

// OnName replaces makes mock to call fn when Symbol function is called.
func (m *ReadOnlyContractMock) OnName(fn func(*bind.CallOpts) (string, error)) {
	m.nameFn = fn
}

// NameCalled return the number of times the function was called in this instance.
func (m *ReadOnlyContractMock) NameCalled() uint64 {
	return atomic.LoadUint64(&m.nameCalled)
}

// Name implement the ReadOnlyContract interface.
func (m *ReadOnlyContractMock) Name(callOpts *bind.CallOpts) (string, error) {
	atomic.AddUint64(&m.nameCalled, 1)
	if m.nameFn != nil {
		return m.nameFn(callOpts)
	}
	return "", nil
}
