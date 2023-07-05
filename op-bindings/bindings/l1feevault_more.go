// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/sliceledger-blockchain/slice-ledger/op-bindings/solc"
)

const L1FeeVaultStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L2/L1FeeVault.sol:L1FeeVault\",\"label\":\"totalProcessed\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint256\"}],\"types\":{\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var L1FeeVaultStorageLayout = new(solc.StorageLayout)

var L1FeeVaultDeployedBin = "0x6080604052600436106100695760003560e01c806384411d651161004357806384411d651461010c578063d0e12f9014610130578063d3e5792b1461017157600080fd5b80630d9019e1146100755780633ccfd60b146100d357806354fd4d50146100ea57600080fd5b3661007057005b600080fd5b34801561008157600080fd5b506100a97f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156100df57600080fd5b506100e86101a5565b005b3480156100f657600080fd5b506100ff6104ab565b6040516100ca919061071b565b34801561011857600080fd5b5061012260005481565b6040519081526020016100ca565b34801561013c57600080fd5b506101647f000000000000000000000000000000000000000000000000000000000000000081565b6040516100ca919061079f565b34801561017d57600080fd5b506101227f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000047101561027f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f4665655661756c743a207769746864726177616c20616d6f756e74206d75737460448201527f2062652067726561746572207468616e206d696e696d756d207769746864726160648201527f77616c20616d6f756e7400000000000000000000000000000000000000000000608482015260a40160405180910390fd5b60004790508060008082825461029591906107e2565b9091555050604080518281527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166020820152338183015290517fc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba9181900360600190a17f38e04cbeb8c10f8f568618aa75be0f10b6729b8b4237743b4de20cbcde2839ee817f0000000000000000000000000000000000000000000000000000000000000000337f000000000000000000000000000000000000000000000000000000000000000060405161038394939291906107fa565b60405180910390a160017f000000000000000000000000000000000000000000000000000000000000000060018111156103bf576103bf610735565b036103f3576103ef7f00000000000000000000000000000000000000000000000000000000000000005a8361054e565b5050565b604080516020810182526000815290517fe11013dd0000000000000000000000000000000000000000000000000000000081527342000000000000000000000000000000000000109163e11013dd918491610476917f0000000000000000000000000000000000000000000000000000000000000000916188b89160040161083b565b6000604051808303818588803b15801561048f57600080fd5b505af11580156104a3573d6000803e3d6000fd5b505050505050565b60606104d67f0000000000000000000000000000000000000000000000000000000000000000610564565b6104ff7f0000000000000000000000000000000000000000000000000000000000000000610564565b6105287f0000000000000000000000000000000000000000000000000000000000000000610564565b60405160200161053a93929190610876565b604051602081830303815290604052905090565b600080600080600080868989f195945050505050565b6060816000036105a757505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b81156105d157806105bb816108ec565b91506105ca9050600a83610953565b91506105ab565b60008167ffffffffffffffff8111156105ec576105ec610967565b6040519080825280601f01601f191660200182016040528015610616576020820181803683370190505b5090505b84156106995761062b600183610996565b9150610638600a866109ad565b6106439060306107e2565b60f81b818381518110610658576106586109c1565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610692600a86610953565b945061061a565b949350505050565b60005b838110156106bc5781810151838201526020016106a4565b838111156106cb576000848401525b50505050565b600081518084526106e98160208601602086016106a1565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061072e60208301846106d1565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6002811061079b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b602081016107ad8284610764565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082198211156107f5576107f56107b3565b500190565b84815273ffffffffffffffffffffffffffffffffffffffff848116602083015283166040820152608081016108326060830184610764565b95945050505050565b73ffffffffffffffffffffffffffffffffffffffff8416815263ffffffff8316602082015260606040820152600061083260608301846106d1565b600084516108888184602089016106a1565b80830190507f2e0000000000000000000000000000000000000000000000000000000000000080825285516108c4816001850160208a016106a1565b600192019182015283516108df8160028401602088016106a1565b0160020195945050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361091d5761091d6107b3565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60008261096257610962610924565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000828210156109a8576109a86107b3565b500390565b6000826109bc576109bc610924565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L1FeeVaultStorageLayoutJSON), L1FeeVaultStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1FeeVault"] = L1FeeVaultStorageLayout
	deployedBytecodes["L1FeeVault"] = L1FeeVaultDeployedBin
}
