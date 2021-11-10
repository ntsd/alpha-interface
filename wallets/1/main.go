package main

import (
	"fmt"

	"github.com/iotaledger/goshimmer/client/wallet/packages/seed"
	"github.com/iotaledger/goshimmer/packages/ledgerstate"
	"github.com/iotaledger/hive.go/crypto/ed25519"
	"github.com/iotaledger/wasp/client"
	"github.com/iotaledger/wasp/client/chainclient"
	"github.com/iotaledger/wasp/client/goshimmer"
	"github.com/iotaledger/wasp/client/scclient"
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/kv/dict"
	"github.com/mr-tron/base58"
)

const chainID = "sX2BhrkneaQiSFKdyv4JA7s5Luoxbuq75TRiddxuUBoP"
const seedBase58 = "DGK8S2NSYrEBze7oYFwrKfVAY2pGhXKxnEKdJ6qHu2Jx"
const smartContractName = "alphainterface"
const waspClientURL = "127.0.0.1:9090"
const goShimmerURL = "https://api.goshimmer.sc.iota.org"

type Wallet struct {
	seed *seed.Seed
}

var addressIndex int

func (w *Wallet) KeyPair() *ed25519.KeyPair {
	return w.seed.KeyPair(uint64(addressIndex))
}

func (w *Wallet) Address() ledgerstate.Address {
	return w.seed.Address(uint64(addressIndex)).Address()
}

func GetCurrentChainID() *iscp.ChainID {
	chid, err := iscp.ChainIDFromBase58(chainID)
	if err != nil {
		panic(err)
	}
	return chid
}

func Client() *chainclient.Client {
	fmt.Println("44")
	goshimmerClient := goshimmer.NewClient(goShimmerURL, 0)
	fmt.Println("55")
	seedBytes, err := base58.Decode(seedBase58)
	if err != nil {
		panic(err)
	}
	wallet := &Wallet{seed.NewSeed(seedBytes)}
	return chainclient.New(
		goshimmerClient,
		client.NewWaspClient(waspClientURL),
		GetCurrentChainID(),
		wallet.KeyPair(),
	)
}

func main() {
	d := dict.New()
	scClient := scclient.New(Client(), iscp.Hn(smartContractName))
	fmt.Println("66")
	r, err := scClient.CallView("getCrops", d)
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}
