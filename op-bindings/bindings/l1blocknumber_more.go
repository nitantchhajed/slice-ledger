// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/sliceledger-blockchain/slice-ledger/op-bindings/solc"
)

const L1BlockNumberStorageLayoutJSON = "{\"storage\":null,\"types\":{}}"

var L1BlockNumberStorageLayout = new(solc.StorageLayout)

var L1BlockNumberDeployedBin = "0x60806040526004361061002d5760003560e01c806354fd4d5014610052578063b9b3efe91461007d57610048565b3661004857600061003c6100a0565b90508060005260206000f35b600061003c6100a0565b34801561005e57600080fd5b50610067610134565b6040516100749190610344565b60405180910390f35b34801561008957600080fd5b506100926100a0565b604051908152602001610074565b600073420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16638381f58a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610101573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101259190610395565b67ffffffffffffffff16905090565b606061015f7f00000000000000000000000000000000000000000000000000000000000000006101d7565b6101887f00000000000000000000000000000000000000000000000000000000000000006101d7565b6101b17f00000000000000000000000000000000000000000000000000000000000000006101d7565b6040516020016101c3939291906103c6565b604051602081830303815290604052905090565b60608160000361021a57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115610244578061022e8161046b565b915061023d9050600a836104d2565b915061021e565b60008167ffffffffffffffff81111561025f5761025f6104e6565b6040519080825280601f01601f191660200182016040528015610289576020820181803683370190505b5090505b841561030c5761029e600183610515565b91506102ab600a8661052c565b6102b6906030610540565b60f81b8183815181106102cb576102cb610558565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610305600a866104d2565b945061028d565b949350505050565b60005b8381101561032f578181015183820152602001610317565b8381111561033e576000848401525b50505050565b6020815260008251806020840152610363816040850160208701610314565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b6000602082840312156103a757600080fd5b815167ffffffffffffffff811681146103bf57600080fd5b9392505050565b600084516103d8818460208901610314565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551610414816001850160208a01610314565b6001920191820152835161042f816002840160208801610314565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361049c5761049c61043c565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000826104e1576104e16104a3565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000828210156105275761052761043c565b500390565b60008261053b5761053b6104a3565b500690565b600082198211156105535761055361043c565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L1BlockNumberStorageLayoutJSON), L1BlockNumberStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1BlockNumber"] = L1BlockNumberStorageLayout
	deployedBytecodes["L1BlockNumber"] = L1BlockNumberDeployedBin
}
