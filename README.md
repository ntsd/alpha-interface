# Alpha Interface

A Hackathon project created by Alpha Interface team for Agri-D Food Hack

## Installation

### Downloading Wasp and wasp-cli

<https://wiki.iota.org/wasp/guide/chains_and_nodes/running-a-node#download-wasp>

### Run A Wasp node

Run A Wasp node will connecting to goshimmer using TXStream.

you can find testnet endpoint here <https://wiki.iota.org/wasp/guide/chains_and_nodes/testnet#endpoints>

for now we using `goshimmer.sc.iota.org:5000`

```Shell
cd wasp-node
wasp
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
wasp-cli peering info -i 0
wasp-cli peering trust {PubKey} 127.0.0.1:4000
wasp-cli peering trust {PubKey} 127.0.0.1:4001
wasp-cli peering list-trusted

# Deploy The Chain
# `committee` will correspond to wasp.0, wasp.1 in `wasp-cli.json`
# `quorum` is minimum amount node
wasp-cli chain deploy --committee=0 --quorum=1 --chain=alpha-interface --description="Alpha Interface Chain"

# trust chain you can find chain id from `http://127.0.0.1:7000/chains`
wasp-cli set chains.testchain {chain-id}
wasp-cli set chain testchain

# Deposit money to the chain
wasp-cli chain deposit IOTA:20000
```

## Resources

<https://wiki.iota.org/wasp/overview>
