// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package example

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// MainABI is the input ABI used to generate the binding from.
const MainABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"d\",\"outputs\":[{\"name\":\"one\",\"type\":\"bytes32\"},{\"name\":\"two\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBoth\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_main\",\"type\":\"bytes32\"},{\"name\":\"_sub\",\"type\":\"bytes32\"}],\"name\":\"set\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"}]"

// MainBin is the compiled bytecode used for deploying new contracts.
const MainBin = `0x6060604052341561000f57600080fd5b5b61034d8061001f6000396000f300606060405263ffffffff60e060020a6000350416638a054ac2811461003a57806393affe511461007b578063f71f7a25146100a6575b600080fd5b341561004557600080fd5b61004d6100d1565b60405191825273ffffffffffffffffffffffffffffffffffffffff1660208201526040908101905180910390f35b341561008657600080fd5b61008e6100f0565b60405191825260208201526040908101905180910390f35b34156100b157600080fd5b6100bf60043560243561016f565b60405190815260200160405180910390f35b60005460015473ffffffffffffffffffffffffffffffffffffffff1682565b6000805460015482919073ffffffffffffffffffffffffffffffffffffffff16636d4ce63c83604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561014b57600080fd5b6102c65a03f1151561015c57600080fd5b50505060405180519050915091505b9091565b60008061017a61023b565b604051809103906000f080151561019057600080fd5b90508073ffffffffffffffffffffffffffffffffffffffff1663db80813f8460405160e060020a63ffffffff84160281526004810191909152602401600060405180830381600087803b15156101e557600080fd5b6102c65a03f115156101f657600080fd5b50505060008490556001805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff83161790555b5092915050565b60405160d78061024b8339019056006060604052341561000f57600080fd5b5b60b98061001e6000396000f300606060405263ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416636d4ce63c81146046578063db80813f146068575b600080fd5b3415605057600080fd5b6056607d565b60405190815260200160405180910390f35b3415607257600080fd5b607b6004356084565b005b6000545b90565b60008190555b505600a165627a7a7230582003cfb96ce9e0cd193e48ea247b0f1b278aee3d998b81cfa05db74fe3c19539ba0029a165627a7a72305820d6979e81f379e5846325d9c561bc6431684fab5b71f9732ade66332b546aabb80029`

// DeployMain deploys a new Ethereum contract, binding an instance of Main to it.
func DeployMain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Main, error) {
	parsed, err := abi.JSON(strings.NewReader(MainABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}}, nil
}

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// D is a free data retrieval call binding the contract method 0x8a054ac2.
//
// Solidity: function d() constant returns(one bytes32, two address)
func (_Main *MainCaller) D(opts *bind.CallOpts) (struct {
	One [32]byte
	Two common.Address
}, error) {
	ret := new(struct {
		One [32]byte
		Two common.Address
	})
	out := ret
	err := _Main.contract.Call(opts, out, "d")
	return *ret, err
}

// D is a free data retrieval call binding the contract method 0x8a054ac2.
//
// Solidity: function d() constant returns(one bytes32, two address)
func (_Main *MainSession) D() (struct {
	One [32]byte
	Two common.Address
}, error) {
	return _Main.Contract.D(&_Main.CallOpts)
}

// D is a free data retrieval call binding the contract method 0x8a054ac2.
//
// Solidity: function d() constant returns(one bytes32, two address)
func (_Main *MainCallerSession) D() (struct {
	One [32]byte
	Two common.Address
}, error) {
	return _Main.Contract.D(&_Main.CallOpts)
}

// GetBoth is a free data retrieval call binding the contract method 0x93affe51.
//
// Solidity: function getBoth() constant returns(bytes32, bytes32)
func (_Main *MainCaller) GetBoth(opts *bind.CallOpts) ([32]byte, [32]byte, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Main.contract.Call(opts, out, "getBoth")
	return *ret0, *ret1, err
}

// GetBoth is a free data retrieval call binding the contract method 0x93affe51.
//
// Solidity: function getBoth() constant returns(bytes32, bytes32)
func (_Main *MainSession) GetBoth() ([32]byte, [32]byte, error) {
	return _Main.Contract.GetBoth(&_Main.CallOpts)
}

// GetBoth is a free data retrieval call binding the contract method 0x93affe51.
//
// Solidity: function getBoth() constant returns(bytes32, bytes32)
func (_Main *MainCallerSession) GetBoth() ([32]byte, [32]byte, error) {
	return _Main.Contract.GetBoth(&_Main.CallOpts)
}

// Set is a paid mutator transaction binding the contract method 0xf71f7a25.
//
// Solidity: function set(_main bytes32, _sub bytes32) returns(uint256)
func (_Main *MainTransactor) Set(opts *bind.TransactOpts, _main [32]byte, _sub [32]byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "set", _main, _sub)
}

// Set is a paid mutator transaction binding the contract method 0xf71f7a25.
//
// Solidity: function set(_main bytes32, _sub bytes32) returns(uint256)
func (_Main *MainSession) Set(_main [32]byte, _sub [32]byte) (*types.Transaction, error) {
	return _Main.Contract.Set(&_Main.TransactOpts, _main, _sub)
}

// Set is a paid mutator transaction binding the contract method 0xf71f7a25.
//
// Solidity: function set(_main bytes32, _sub bytes32) returns(uint256)
func (_Main *MainTransactorSession) Set(_main [32]byte, _sub [32]byte) (*types.Transaction, error) {
	return _Main.Contract.Set(&_Main.TransactOpts, _main, _sub)
}

// SubABI is the input ABI used to generate the binding from.
const SubABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_one\",\"type\":\"bytes32\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"type\":\"function\"}]"

// SubBin is the compiled bytecode used for deploying new contracts.
const SubBin = `0x6060604052341561000f57600080fd5b5b60b98061001e6000396000f300606060405263ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416636d4ce63c81146046578063db80813f146068575b600080fd5b3415605057600080fd5b6056607d565b60405190815260200160405180910390f35b3415607257600080fd5b607b6004356084565b005b6000545b90565b60008190555b505600a165627a7a7230582003cfb96ce9e0cd193e48ea247b0f1b278aee3d998b81cfa05db74fe3c19539ba0029`

// DeploySub deploys a new Ethereum contract, binding an instance of Sub to it.
func DeploySub(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Sub, error) {
	parsed, err := abi.JSON(strings.NewReader(SubABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SubBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sub{SubCaller: SubCaller{contract: contract}, SubTransactor: SubTransactor{contract: contract}}, nil
}

// Sub is an auto generated Go binding around an Ethereum contract.
type Sub struct {
	SubCaller     // Read-only binding to the contract
	SubTransactor // Write-only binding to the contract
}

// SubCaller is an auto generated read-only Go binding around an Ethereum contract.
type SubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SubSession struct {
	Contract     *Sub              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SubCallerSession struct {
	Contract *SubCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SubTransactorSession struct {
	Contract     *SubTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SubRaw is an auto generated low-level Go binding around an Ethereum contract.
type SubRaw struct {
	Contract *Sub // Generic contract binding to access the raw methods on
}

// SubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SubCallerRaw struct {
	Contract *SubCaller // Generic read-only contract binding to access the raw methods on
}

// SubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SubTransactorRaw struct {
	Contract *SubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSub creates a new instance of Sub, bound to a specific deployed contract.
func NewSub(address common.Address, backend bind.ContractBackend) (*Sub, error) {
	contract, err := bindSub(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sub{SubCaller: SubCaller{contract: contract}, SubTransactor: SubTransactor{contract: contract}}, nil
}

// NewSubCaller creates a new read-only instance of Sub, bound to a specific deployed contract.
func NewSubCaller(address common.Address, caller bind.ContractCaller) (*SubCaller, error) {
	contract, err := bindSub(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SubCaller{contract: contract}, nil
}

// NewSubTransactor creates a new write-only instance of Sub, bound to a specific deployed contract.
func NewSubTransactor(address common.Address, transactor bind.ContractTransactor) (*SubTransactor, error) {
	contract, err := bindSub(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SubTransactor{contract: contract}, nil
}

// bindSub binds a generic wrapper to an already deployed contract.
func bindSub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sub *SubRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Sub.Contract.SubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sub *SubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sub.Contract.SubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sub *SubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sub.Contract.SubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sub *SubCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Sub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sub *SubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sub *SubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sub.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() constant returns(bytes32)
func (_Sub *SubCaller) Get(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Sub.contract.Call(opts, out, "get")
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() constant returns(bytes32)
func (_Sub *SubSession) Get() ([32]byte, error) {
	return _Sub.Contract.Get(&_Sub.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() constant returns(bytes32)
func (_Sub *SubCallerSession) Get() ([32]byte, error) {
	return _Sub.Contract.Get(&_Sub.CallOpts)
}

// Set is a paid mutator transaction binding the contract method 0xdb80813f.
//
// Solidity: function set(_one bytes32) returns()
func (_Sub *SubTransactor) Set(opts *bind.TransactOpts, _one [32]byte) (*types.Transaction, error) {
	return _Sub.contract.Transact(opts, "set", _one)
}

// Set is a paid mutator transaction binding the contract method 0xdb80813f.
//
// Solidity: function set(_one bytes32) returns()
func (_Sub *SubSession) Set(_one [32]byte) (*types.Transaction, error) {
	return _Sub.Contract.Set(&_Sub.TransactOpts, _one)
}

// Set is a paid mutator transaction binding the contract method 0xdb80813f.
//
// Solidity: function set(_one bytes32) returns()
func (_Sub *SubTransactorSession) Set(_one [32]byte) (*types.Transaction, error) {
	return _Sub.Contract.Set(&_Sub.TransactOpts, _one)
}
