package state_test

import (
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/sliceledger-blockchain/slice-ledger/op-chain-ops/state"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestAddBalance(t *testing.T) {
	t.Parallel()

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	db := state.NewMemoryStateDB(nil)

	for i := 0; i < 100; i++ {
		key, _ := crypto.GenerateKey()
		addr := crypto.PubkeyToAddress(key.PublicKey)
		value := new(big.Int).Rand(rng, big.NewInt(1000))

		db.CreateAccount(addr)
		db.AddBalance(addr, value)

		account := db.GetAccount(addr)
		require.NotNil(t, account)
		require.True(t, BigEqual(account.Balance, value))
	}
}

func TestCode(t *testing.T) {
	t.Parallel()

	db := state.NewMemoryStateDB(nil)

	for i := 0; i < 100; i++ {
		key, _ := crypto.GenerateKey()
		addr := crypto.PubkeyToAddress(key.PublicKey)

		db.CreateAccount(addr)

		pre := db.GetCode(addr)
		require.Nil(t, pre)

		code := make([]byte, rand.Intn(1024))
		rand.Read(code)

		db.SetCode(addr, code)

		post := db.GetCode(addr)
		require.Equal(t, post, code)

		size := db.GetCodeSize(addr)
		require.Equal(t, size, len(code))

		codeHash := db.GetCodeHash(addr)
		require.Equal(t, codeHash, common.BytesToHash(crypto.Keccak256(code)))
	}
}

func BigEqual(a, b *big.Int) bool {
	if a == nil || b == nil {
		return a == b
	} else {
		return a.Cmp(b) == 0
	}
}
