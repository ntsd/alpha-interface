# Alpha Interface

A Hackathon project created by Alpha Interface team for Agri-D Food Hack

## Installation

### Downloading Wasp and wasp-cli

<https://wiki.iota.org/wasp/guide/chains_and_nodes/running-a-node#download-wasp>

```SHELL
brew install rocksdb
make install
```

### Run A Wasp node

Run A Wasp node will connecting to goshimmer using TXStream.

you can find testnet endpoint here <https://wiki.iota.org/wasp/guide/chains_and_nodes/testnet#endpoints>

for now we using `goshimmer.sc.iota.org:5000`

```Shell
wasp -c ./wasp-nodes/config.json

# or

make run-wasp
```

### Configuring wasp-cli

```shell
wasp-cli init

# Set go shimmer api and request fund
wasp-cli set goshimmer.api https://api.goshimmer.sc.iota.org
wasp-cli request-funds

# set wasp address for a local node
wasp-cli set wasp.0.api 127.0.0.1:9090
wasp-cli set wasp.0.nanomsg 127.0.0.1:5550
wasp-cli set wasp.0.peering 127.0.0.1:4000
```

### Setting Up a Chain

```Shell
# Trust node
wasp-cli peering info
wasp-cli peering trust {PubKey} 127.0.0.1:4000
wasp-cli peering list-trusted

# Deploy The Chain
# `committee` will correspond to wasp.0, wasp.1 in `wasp-cli.json`
# `quorum` is minimum amount node
wasp-cli chain deploy --committee=0 --quorum=1 --chain=alpha-interface-chain --description="Alpha Interface Chain"

# Deposit money to the chain
wasp-cli chain deposit IOTA:10000

# Set test chain you can find chain id from `http://127.0.0.1:7000/chains`
wasp-cli set chains.testchain {chain-id}
wasp-cli set chain testchain
```

### Build the smart contract

```shell
make build-wasm
```

### Deploy the smart contract

```Shell
make deploy-wasm
```

### Post smart contract request

```Shell
wasp-cli chain post-request alphainterface <funcname> [params] --transfer=IOTA:10
wasp-cli chain post-request alphainterface setOwner string owner string {actorID}
wasp-cli chain call-view alphainterface getOwner
wasp-cli chain post-request alphainterface setCrop string name string potato string country string germany string yield int 10000 --off-ledger
wasp-cli chain post-request alphainterface setCrop string name string rice string country string germany string yield int 20000 --off-ledger
wasp-cli chain call-view alphainterface getCrops
wasp-cli chain call-view alphainterface getCrop string cropIdx int32 0
```

## Resources

<https://wiki.iota.org/wasp/overview>
