package example

//go:generate abigen --sol ./example.sol --pkg example --out ./example.go

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	clientKey, _ = crypto.HexToECDSA("49a7b37aa6f6645917e7b807e9d1c00d4fa71f18343b0d4122a4d2df64dd6fee")
	clientAddr   = crypto.PubkeyToAddress(clientKey.PublicKey)
)

func newTestBackend() *backends.SimulatedBackend {
	return backends.NewSimulatedBackend(core.GenesisAlloc{
		clientAddr: {Balance: big.NewInt(1000000000)},
	})
}

//commit new block to the backends.SimulatedBackend
//the blockcount is increment on each commit.
func commit(b *backends.SimulatedBackend) {
	b.Commit()
}

func TestExample(t *testing.T) {
	backend := newTestBackend()
	deployTransactor := bind.NewKeyedTransactor(clientKey)
	deployTransactor.Value = big.NewInt(0)
	add, main, err := deployExample(clientKey, big.NewInt(0), backend)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("Main Address", add.String())
	commit(backend)

	_, err = main.Set(deployTransactor, [32]byte{1, 2, 4}, [32]byte{4, 5, 6})
	if err != nil {
		t.Fatal(err)
	}
	commit(backend)

	_main, _sub, err := main.GetBoth(&bind.CallOpts{})

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("main", _main, "sub", _sub)

}
func deployExample(prvKey *ecdsa.PrivateKey, amount *big.Int, backend *backends.SimulatedBackend) (common.Address, *Main, error) {
	deployTransactor := bind.NewKeyedTransactor(prvKey)
	deployTransactor.Value = amount

	addr, _, mirror, err := DeployMain(deployTransactor, backend)
	if err != nil {
		return common.Address{}, nil, err
	}
	commit(backend)
	return addr, mirror, nil
}
