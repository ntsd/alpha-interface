package main

import (
	"fmt"

	"github.com/iotaledger/wasp/client"
	"github.com/iotaledger/wasp/client/chainclient"
	"github.com/iotaledger/wasp/client/goshimmer"
	"github.com/iotaledger/wasp/client/scclient"
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/kv/dict"
	"github.com/iotaledger/wasp/tools/wasp-cli/log"
	"github.com/iotaledger/wasp/tools/wasp-cli/wallet"
)

const chain = "ecee4ZYGG4YcqWfcwHmQTzXJerUXFweELgAzoPU5sUdB"
const smartContract = "alphainterface"
const waspClient = "127.0.0.1:9090"
const goShimmer = "https://api.goshimmer.sc.iota.org"

func GetCurrentChainID() *iscp.ChainID {
	chid, err := iscp.ChainIDFromBase58(chain)
	log.Check(err)
	return chid
}

func Client() *chainclient.Client {
	goshimmerClient := goshimmer.NewClient(goShimmer, 0)
	return chainclient.New(
		goshimmerClient,
		client.NewWaspClient(waspClient),
		GetCurrentChainID(),
		wallet.Load().KeyPair(),
	)
}

func main() {
	d := dict.New()
	r, err := scclient.New(Client(), iscp.Hn(smartContract)).CallView("getCrops", d)
	log.Check(err)
	fmt.Println(r)
}
