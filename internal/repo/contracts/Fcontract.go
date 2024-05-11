// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AttendanceContractAttendanceRecord is an auto generated low-level Go binding around an user-defined struct.
type AttendanceContractAttendanceRecord struct {
	EmployeeId  *big.Int
	CheckInTime *big.Int
	Details     string
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"employeeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"checkInTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"name\":\"AttendanceRecorded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"employeeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCheckInTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newDetails\",\"type\":\"string\"}],\"name\":\"AttendanceUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_entity\",\"type\":\"address\"}],\"name\":\"authorizeEntity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_employeeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_end\",\"type\":\"uint256\"}],\"name\":\"getAttendanceByDateRange\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"employeeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"checkInTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"internalType\":\"structAttendanceContract.AttendanceRecord[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_employeeId\",\"type\":\"uint256\"}],\"name\":\"getAttendanceByEmployeeId\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"employeeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"checkInTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"internalType\":\"structAttendanceContract.AttendanceRecord[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEmployeeID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_emID\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_employeeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_checkInTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_details\",\"type\":\"string\"}],\"name\":\"recordAttendance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_entity\",\"type\":\"address\"}],\"name\":\"revokeAuthorization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_employeeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newCheckInTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_newDetails\",\"type\":\"string\"}],\"name\":\"updateAttendance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260015f553480156012575f80fd5b50600160025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208190555061153c806100635f395ff3fe608060405234801561000f575f80fd5b506004361061007b575f3560e01c80639f43f46f116100595780639f43f46f146100d5578063b48028e314610105578063b608d3a414610121578063dd0826c61461013d5761007b565b8063020310241461007f57806308c220c21461009b5780630c6b5348146100b7575b5f80fd5b61009960048036038101906100949190610b5a565b61016d565b005b6100b560048036038101906100b09190610cf4565b61024c565b005b6100bf610418565b6040516100cc9190610d6f565b60405180910390f35b6100ef60048036038101906100ea9190610d88565b61055a565b6040516100fc9190610f4f565b60405180910390f35b61011f600480360381019061011a9190610b5a565b61076e565b005b61013b60048036038101906101369190610f6f565b610832565b005b61015760048036038101906101529190610fef565b6109c0565b6040516101649190610f4f565b60405180910390f35b5f60025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054036101ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101e390611074565b60405180910390fd5b5f808154809291906101fd906110bf565b919050555f819055505f5460025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208190555050565b5f60025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054036102cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102c290611074565b60405180910390fd5b620f424060025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205414158015610358575060025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20548314155b6104135760015f8481526020019081526020015f20604051806060016040528085815260200184815260200183815250908060018154018082558091505060019003905f5260205f2090600302015f909190919091505f820151815f01556020820151816001015560408201518160020190816103d59190611300565b505050827f32e7c28972e89757b6baa3898e54af0a57a00f3a4df3d1ad95609eb202460b3a838360405161040a929190611407565b60405180910390a25b505050565b5f8060025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205403610498576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161048f90611074565b60405180910390fd5b5f60025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205403610517576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050e90611074565b60405180910390fd5b60025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905090565b60605f60015f8681526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b82821015610660578382905f5260205f2090600302016040518060600160405290815f8201548152602001600182015481526020016002820180546105d190611133565b80601f01602080910402602001604051908101604052809291908181526020018280546105fd90611133565b80156106485780601f1061061f57610100808354040283529160200191610648565b820191905f5260205f20905b81548152906001019060200180831161062b57829003601f168201915b5050505050815250508152602001906001019061058d565b5050505090505f815167ffffffffffffffff81111561068257610681610bd0565b5b6040519080825280602002602001820160405280156106bb57816020015b6106a8610ad0565b8152602001906001900390816106a05790505b5090505f5b825181101561076157858382815181106106dd576106dc611435565b5b6020026020010151602001511015801561071557508483828151811061070657610705611435565b5b60200260200101516020015111155b156107545782818151811061072d5761072c611435565b5b602002602001015182828151811061074857610747611435565b5b60200260200101819052505b80806001019150506106c0565b5080925050509392505050565b5f60025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054036107ed576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107e490611074565b60405180910390fd5b5f60025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208190555050565b5f60025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054036108b1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108a890611074565b60405180910390fd5b60015f8581526020019081526020015f20805490508310610907576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108fe906114ac565b60405180910390fd5b8160015f8681526020019081526020015f20848154811061092b5761092a611435565b5b905f5260205f209060030201600101819055508060015f8681526020019081526020015f20848154811061096257610961611435565b5b905f5260205f209060030201600201908161097d9190611300565b50837fa0d3071670661a556feb29023709b20ab2a62d5c56f63343d67c087d28be3df78484846040516109b2939291906114ca565b60405180910390a250505050565b606060015f8381526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b82821015610ac5578382905f5260205f2090600302016040518060600160405290815f820154815260200160018201548152602001600282018054610a3690611133565b80601f0160208091040260200160405190810160405280929190818152602001828054610a6290611133565b8015610aad5780601f10610a8457610100808354040283529160200191610aad565b820191905f5260205f20905b815481529060010190602001808311610a9057829003601f168201915b505050505081525050815260200190600101906109f2565b505050509050919050565b60405180606001604052805f81526020015f8152602001606081525090565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610b2982610b00565b9050919050565b610b3981610b1f565b8114610b43575f80fd5b50565b5f81359050610b5481610b30565b92915050565b5f60208284031215610b6f57610b6e610af8565b5b5f610b7c84828501610b46565b91505092915050565b5f819050919050565b610b9781610b85565b8114610ba1575f80fd5b50565b5f81359050610bb281610b8e565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610c0682610bc0565b810181811067ffffffffffffffff82111715610c2557610c24610bd0565b5b80604052505050565b5f610c37610aef565b9050610c438282610bfd565b919050565b5f67ffffffffffffffff821115610c6257610c61610bd0565b5b610c6b82610bc0565b9050602081019050919050565b828183375f83830152505050565b5f610c98610c9384610c48565b610c2e565b905082815260208101848484011115610cb457610cb3610bbc565b5b610cbf848285610c78565b509392505050565b5f82601f830112610cdb57610cda610bb8565b5b8135610ceb848260208601610c86565b91505092915050565b5f805f60608486031215610d0b57610d0a610af8565b5b5f610d1886828701610ba4565b9350506020610d2986828701610ba4565b925050604084013567ffffffffffffffff811115610d4a57610d49610afc565b5b610d5686828701610cc7565b9150509250925092565b610d6981610b85565b82525050565b5f602082019050610d825f830184610d60565b92915050565b5f805f60608486031215610d9f57610d9e610af8565b5b5f610dac86828701610ba4565b9350506020610dbd86828701610ba4565b9250506040610dce86828701610ba4565b9150509250925092565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610e0a81610b85565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f610e4282610e10565b610e4c8185610e1a565b9350610e5c818560208601610e2a565b610e6581610bc0565b840191505092915050565b5f606083015f830151610e855f860182610e01565b506020830151610e986020860182610e01565b5060408301518482036040860152610eb08282610e38565b9150508091505092915050565b5f610ec88383610e70565b905092915050565b5f602082019050919050565b5f610ee682610dd8565b610ef08185610de2565b935083602082028501610f0285610df2565b805f5b85811015610f3d5784840389528151610f1e8582610ebd565b9450610f2983610ed0565b925060208a01995050600181019050610f05565b50829750879550505050505092915050565b5f6020820190508181035f830152610f678184610edc565b905092915050565b5f805f8060808587031215610f8757610f86610af8565b5b5f610f9487828801610ba4565b9450506020610fa587828801610ba4565b9350506040610fb687828801610ba4565b925050606085013567ffffffffffffffff811115610fd757610fd6610afc565b5b610fe387828801610cc7565b91505092959194509250565b5f6020828403121561100457611003610af8565b5b5f61101184828501610ba4565b91505092915050565b5f82825260208201905092915050565b7f556e617574686f72697a656420616363657373000000000000000000000000005f82015250565b5f61105e60138361101a565b91506110698261102a565b602082019050919050565b5f6020820190508181035f83015261108b81611052565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6110c982610b85565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036110fb576110fa611092565b5b600182019050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061114a57607f821691505b60208210810361115d5761115c611106565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026111bf7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611184565b6111c98683611184565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6112046111ff6111fa84610b85565b6111e1565b610b85565b9050919050565b5f819050919050565b61121d836111ea565b6112316112298261120b565b848454611190565b825550505050565b5f90565b611245611239565b611250818484611214565b505050565b5b81811015611273576112685f8261123d565b600181019050611256565b5050565b601f8211156112b85761128981611163565b61129284611175565b810160208510156112a1578190505b6112b56112ad85611175565b830182611255565b50505b505050565b5f82821c905092915050565b5f6112d85f19846008026112bd565b1980831691505092915050565b5f6112f083836112c9565b9150826002028217905092915050565b61130982610e10565b67ffffffffffffffff81111561132257611321610bd0565b5b61132c8254611133565b611337828285611277565b5f60209050601f831160018114611368575f8415611356578287015190505b61136085826112e5565b8655506113c7565b601f19841661137686611163565b5f5b8281101561139d57848901518255600182019150602085019450602081019050611378565b868310156113ba57848901516113b6601f8916826112c9565b8355505b6001600288020188555050505b505050505050565b5f6113d982610e10565b6113e3818561101a565b93506113f3818560208601610e2a565b6113fc81610bc0565b840191505092915050565b5f60408201905061141a5f830185610d60565b818103602083015261142c81846113cf565b90509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f496e646578206f7574206f6620626f756e6473000000000000000000000000005f82015250565b5f61149660138361101a565b91506114a182611462565b602082019050919050565b5f6020820190508181035f8301526114c38161148a565b9050919050565b5f6060820190506114dd5f830186610d60565b6114ea6020830185610d60565b81810360408301526114fc81846113cf565b905094935050505056fea264697066735822122043997a0e10f250205a699955abb09418c469f4f2e425d74f9097abba850131bd64736f6c63430008190033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// GetAttendanceByDateRange is a free data retrieval call binding the contract method 0x9f43f46f.
//
// Solidity: function getAttendanceByDateRange(uint256 _employeeId, uint256 _start, uint256 _end) view returns((uint256,uint256,string)[])
func (_Contracts *ContractsCaller) GetAttendanceByDateRange(opts *bind.CallOpts, _employeeId *big.Int, _start *big.Int, _end *big.Int) ([]AttendanceContractAttendanceRecord, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAttendanceByDateRange", _employeeId, _start, _end)

	if err != nil {
		return *new([]AttendanceContractAttendanceRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]AttendanceContractAttendanceRecord)).(*[]AttendanceContractAttendanceRecord)

	return out0, err

}

// GetAttendanceByDateRange is a free data retrieval call binding the contract method 0x9f43f46f.
//
// Solidity: function getAttendanceByDateRange(uint256 _employeeId, uint256 _start, uint256 _end) view returns((uint256,uint256,string)[])
func (_Contracts *ContractsSession) GetAttendanceByDateRange(_employeeId *big.Int, _start *big.Int, _end *big.Int) ([]AttendanceContractAttendanceRecord, error) {
	return _Contracts.Contract.GetAttendanceByDateRange(&_Contracts.CallOpts, _employeeId, _start, _end)
}

// GetAttendanceByDateRange is a free data retrieval call binding the contract method 0x9f43f46f.
//
// Solidity: function getAttendanceByDateRange(uint256 _employeeId, uint256 _start, uint256 _end) view returns((uint256,uint256,string)[])
func (_Contracts *ContractsCallerSession) GetAttendanceByDateRange(_employeeId *big.Int, _start *big.Int, _end *big.Int) ([]AttendanceContractAttendanceRecord, error) {
	return _Contracts.Contract.GetAttendanceByDateRange(&_Contracts.CallOpts, _employeeId, _start, _end)
}

// GetAttendanceByEmployeeId is a free data retrieval call binding the contract method 0xdd0826c6.
//
// Solidity: function getAttendanceByEmployeeId(uint256 _employeeId) view returns((uint256,uint256,string)[])
func (_Contracts *ContractsCaller) GetAttendanceByEmployeeId(opts *bind.CallOpts, _employeeId *big.Int) ([]AttendanceContractAttendanceRecord, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAttendanceByEmployeeId", _employeeId)

	if err != nil {
		return *new([]AttendanceContractAttendanceRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]AttendanceContractAttendanceRecord)).(*[]AttendanceContractAttendanceRecord)

	return out0, err

}

// GetAttendanceByEmployeeId is a free data retrieval call binding the contract method 0xdd0826c6.
//
// Solidity: function getAttendanceByEmployeeId(uint256 _employeeId) view returns((uint256,uint256,string)[])
func (_Contracts *ContractsSession) GetAttendanceByEmployeeId(_employeeId *big.Int) ([]AttendanceContractAttendanceRecord, error) {
	return _Contracts.Contract.GetAttendanceByEmployeeId(&_Contracts.CallOpts, _employeeId)
}

// GetAttendanceByEmployeeId is a free data retrieval call binding the contract method 0xdd0826c6.
//
// Solidity: function getAttendanceByEmployeeId(uint256 _employeeId) view returns((uint256,uint256,string)[])
func (_Contracts *ContractsCallerSession) GetAttendanceByEmployeeId(_employeeId *big.Int) ([]AttendanceContractAttendanceRecord, error) {
	return _Contracts.Contract.GetAttendanceByEmployeeId(&_Contracts.CallOpts, _employeeId)
}

// GetEmployeeID is a free data retrieval call binding the contract method 0x0c6b5348.
//
// Solidity: function getEmployeeID() view returns(uint256 _emID)
func (_Contracts *ContractsCaller) GetEmployeeID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getEmployeeID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEmployeeID is a free data retrieval call binding the contract method 0x0c6b5348.
//
// Solidity: function getEmployeeID() view returns(uint256 _emID)
func (_Contracts *ContractsSession) GetEmployeeID() (*big.Int, error) {
	return _Contracts.Contract.GetEmployeeID(&_Contracts.CallOpts)
}

// GetEmployeeID is a free data retrieval call binding the contract method 0x0c6b5348.
//
// Solidity: function getEmployeeID() view returns(uint256 _emID)
func (_Contracts *ContractsCallerSession) GetEmployeeID() (*big.Int, error) {
	return _Contracts.Contract.GetEmployeeID(&_Contracts.CallOpts)
}

// AuthorizeEntity is a paid mutator transaction binding the contract method 0x02031024.
//
// Solidity: function authorizeEntity(address _entity) returns()
func (_Contracts *ContractsTransactor) AuthorizeEntity(opts *bind.TransactOpts, _entity common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "authorizeEntity", _entity)
}

// AuthorizeEntity is a paid mutator transaction binding the contract method 0x02031024.
//
// Solidity: function authorizeEntity(address _entity) returns()
func (_Contracts *ContractsSession) AuthorizeEntity(_entity common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AuthorizeEntity(&_Contracts.TransactOpts, _entity)
}

// AuthorizeEntity is a paid mutator transaction binding the contract method 0x02031024.
//
// Solidity: function authorizeEntity(address _entity) returns()
func (_Contracts *ContractsTransactorSession) AuthorizeEntity(_entity common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AuthorizeEntity(&_Contracts.TransactOpts, _entity)
}

// RecordAttendance is a paid mutator transaction binding the contract method 0x08c220c2.
//
// Solidity: function recordAttendance(uint256 _employeeId, uint256 _checkInTime, string _details) returns()
func (_Contracts *ContractsTransactor) RecordAttendance(opts *bind.TransactOpts, _employeeId *big.Int, _checkInTime *big.Int, _details string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "recordAttendance", _employeeId, _checkInTime, _details)
}

// RecordAttendance is a paid mutator transaction binding the contract method 0x08c220c2.
//
// Solidity: function recordAttendance(uint256 _employeeId, uint256 _checkInTime, string _details) returns()
func (_Contracts *ContractsSession) RecordAttendance(_employeeId *big.Int, _checkInTime *big.Int, _details string) (*types.Transaction, error) {
	return _Contracts.Contract.RecordAttendance(&_Contracts.TransactOpts, _employeeId, _checkInTime, _details)
}

// RecordAttendance is a paid mutator transaction binding the contract method 0x08c220c2.
//
// Solidity: function recordAttendance(uint256 _employeeId, uint256 _checkInTime, string _details) returns()
func (_Contracts *ContractsTransactorSession) RecordAttendance(_employeeId *big.Int, _checkInTime *big.Int, _details string) (*types.Transaction, error) {
	return _Contracts.Contract.RecordAttendance(&_Contracts.TransactOpts, _employeeId, _checkInTime, _details)
}

// RevokeAuthorization is a paid mutator transaction binding the contract method 0xb48028e3.
//
// Solidity: function revokeAuthorization(address _entity) returns()
func (_Contracts *ContractsTransactor) RevokeAuthorization(opts *bind.TransactOpts, _entity common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "revokeAuthorization", _entity)
}

// RevokeAuthorization is a paid mutator transaction binding the contract method 0xb48028e3.
//
// Solidity: function revokeAuthorization(address _entity) returns()
func (_Contracts *ContractsSession) RevokeAuthorization(_entity common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RevokeAuthorization(&_Contracts.TransactOpts, _entity)
}

// RevokeAuthorization is a paid mutator transaction binding the contract method 0xb48028e3.
//
// Solidity: function revokeAuthorization(address _entity) returns()
func (_Contracts *ContractsTransactorSession) RevokeAuthorization(_entity common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RevokeAuthorization(&_Contracts.TransactOpts, _entity)
}

// UpdateAttendance is a paid mutator transaction binding the contract method 0xb608d3a4.
//
// Solidity: function updateAttendance(uint256 _employeeId, uint256 _index, uint256 _newCheckInTime, string _newDetails) returns()
func (_Contracts *ContractsTransactor) UpdateAttendance(opts *bind.TransactOpts, _employeeId *big.Int, _index *big.Int, _newCheckInTime *big.Int, _newDetails string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "updateAttendance", _employeeId, _index, _newCheckInTime, _newDetails)
}

// UpdateAttendance is a paid mutator transaction binding the contract method 0xb608d3a4.
//
// Solidity: function updateAttendance(uint256 _employeeId, uint256 _index, uint256 _newCheckInTime, string _newDetails) returns()
func (_Contracts *ContractsSession) UpdateAttendance(_employeeId *big.Int, _index *big.Int, _newCheckInTime *big.Int, _newDetails string) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateAttendance(&_Contracts.TransactOpts, _employeeId, _index, _newCheckInTime, _newDetails)
}

// UpdateAttendance is a paid mutator transaction binding the contract method 0xb608d3a4.
//
// Solidity: function updateAttendance(uint256 _employeeId, uint256 _index, uint256 _newCheckInTime, string _newDetails) returns()
func (_Contracts *ContractsTransactorSession) UpdateAttendance(_employeeId *big.Int, _index *big.Int, _newCheckInTime *big.Int, _newDetails string) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateAttendance(&_Contracts.TransactOpts, _employeeId, _index, _newCheckInTime, _newDetails)
}

// ContractsAttendanceRecordedIterator is returned from FilterAttendanceRecorded and is used to iterate over the raw logs and unpacked data for AttendanceRecorded events raised by the Contracts contract.
type ContractsAttendanceRecordedIterator struct {
	Event *ContractsAttendanceRecorded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsAttendanceRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsAttendanceRecorded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsAttendanceRecorded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsAttendanceRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsAttendanceRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsAttendanceRecorded represents a AttendanceRecorded event raised by the Contracts contract.
type ContractsAttendanceRecorded struct {
	EmployeeId  *big.Int
	CheckInTime *big.Int
	Details     string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAttendanceRecorded is a free log retrieval operation binding the contract event 0x32e7c28972e89757b6baa3898e54af0a57a00f3a4df3d1ad95609eb202460b3a.
//
// Solidity: event AttendanceRecorded(uint256 indexed employeeId, uint256 checkInTime, string details)
func (_Contracts *ContractsFilterer) FilterAttendanceRecorded(opts *bind.FilterOpts, employeeId []*big.Int) (*ContractsAttendanceRecordedIterator, error) {

	var employeeIdRule []interface{}
	for _, employeeIdItem := range employeeId {
		employeeIdRule = append(employeeIdRule, employeeIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "AttendanceRecorded", employeeIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsAttendanceRecordedIterator{contract: _Contracts.contract, event: "AttendanceRecorded", logs: logs, sub: sub}, nil
}

// WatchAttendanceRecorded is a free log subscription operation binding the contract event 0x32e7c28972e89757b6baa3898e54af0a57a00f3a4df3d1ad95609eb202460b3a.
//
// Solidity: event AttendanceRecorded(uint256 indexed employeeId, uint256 checkInTime, string details)
func (_Contracts *ContractsFilterer) WatchAttendanceRecorded(opts *bind.WatchOpts, sink chan<- *ContractsAttendanceRecorded, employeeId []*big.Int) (event.Subscription, error) {

	var employeeIdRule []interface{}
	for _, employeeIdItem := range employeeId {
		employeeIdRule = append(employeeIdRule, employeeIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "AttendanceRecorded", employeeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsAttendanceRecorded)
				if err := _Contracts.contract.UnpackLog(event, "AttendanceRecorded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAttendanceRecorded is a log parse operation binding the contract event 0x32e7c28972e89757b6baa3898e54af0a57a00f3a4df3d1ad95609eb202460b3a.
//
// Solidity: event AttendanceRecorded(uint256 indexed employeeId, uint256 checkInTime, string details)
func (_Contracts *ContractsFilterer) ParseAttendanceRecorded(log types.Log) (*ContractsAttendanceRecorded, error) {
	event := new(ContractsAttendanceRecorded)
	if err := _Contracts.contract.UnpackLog(event, "AttendanceRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsAttendanceUpdatedIterator is returned from FilterAttendanceUpdated and is used to iterate over the raw logs and unpacked data for AttendanceUpdated events raised by the Contracts contract.
type ContractsAttendanceUpdatedIterator struct {
	Event *ContractsAttendanceUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsAttendanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsAttendanceUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsAttendanceUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsAttendanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsAttendanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsAttendanceUpdated represents a AttendanceUpdated event raised by the Contracts contract.
type ContractsAttendanceUpdated struct {
	EmployeeId     *big.Int
	Index          *big.Int
	NewCheckInTime *big.Int
	NewDetails     string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAttendanceUpdated is a free log retrieval operation binding the contract event 0xa0d3071670661a556feb29023709b20ab2a62d5c56f63343d67c087d28be3df7.
//
// Solidity: event AttendanceUpdated(uint256 indexed employeeId, uint256 index, uint256 newCheckInTime, string newDetails)
func (_Contracts *ContractsFilterer) FilterAttendanceUpdated(opts *bind.FilterOpts, employeeId []*big.Int) (*ContractsAttendanceUpdatedIterator, error) {

	var employeeIdRule []interface{}
	for _, employeeIdItem := range employeeId {
		employeeIdRule = append(employeeIdRule, employeeIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "AttendanceUpdated", employeeIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsAttendanceUpdatedIterator{contract: _Contracts.contract, event: "AttendanceUpdated", logs: logs, sub: sub}, nil
}

// WatchAttendanceUpdated is a free log subscription operation binding the contract event 0xa0d3071670661a556feb29023709b20ab2a62d5c56f63343d67c087d28be3df7.
//
// Solidity: event AttendanceUpdated(uint256 indexed employeeId, uint256 index, uint256 newCheckInTime, string newDetails)
func (_Contracts *ContractsFilterer) WatchAttendanceUpdated(opts *bind.WatchOpts, sink chan<- *ContractsAttendanceUpdated, employeeId []*big.Int) (event.Subscription, error) {

	var employeeIdRule []interface{}
	for _, employeeIdItem := range employeeId {
		employeeIdRule = append(employeeIdRule, employeeIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "AttendanceUpdated", employeeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsAttendanceUpdated)
				if err := _Contracts.contract.UnpackLog(event, "AttendanceUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAttendanceUpdated is a log parse operation binding the contract event 0xa0d3071670661a556feb29023709b20ab2a62d5c56f63343d67c087d28be3df7.
//
// Solidity: event AttendanceUpdated(uint256 indexed employeeId, uint256 index, uint256 newCheckInTime, string newDetails)
func (_Contracts *ContractsFilterer) ParseAttendanceUpdated(log types.Log) (*ContractsAttendanceUpdated, error) {
	event := new(ContractsAttendanceUpdated)
	if err := _Contracts.contract.UnpackLog(event, "AttendanceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
