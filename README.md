# Bundle Go SDK


## Disclaimer
**The software and related documentation are under active development, all subject to potential future change without
notification and not ready for production use. The code and security audit have not been fully completed and not ready
for any bug bounty. We advise you to be careful and experiment on the network at your own risk. Stay safe out there.**

## Instruction

The BUNDLE-GO-SDK provides a thin wrapper for interacting with bundle apis on bsc-mainnet & bsc-testnet.

### Requirement

Go version above 1.20

## Getting started
To get started working with the SDK setup your project for Go modules, and retrieve the SDK dependencies with `go get`.
This example shows how you can use the bundle go SDK to interact with the bundle apis on bsc,

### Initialize Project

```sh
$ mkdir ~/hello_bundle
$ cd ~/hello_bundle
$ go mod init hello_bundle
```

### Add SDK Dependencies

```sh
$ go get github.com/node-real/bundle-go-sdk
```

### Initialize Client

The bundle client requires the following parameters to connect to bsc chain.

| Parameter             | Description                                       |
|:----------------------|:--------------------------------------------------|
| rpcAddr               | the tendermit address of greenfield chain         |
| chainId               | the chain id of greenfield                        |
| client.Option  | All the options such as DefaultAccount and secure |

```go
package main

import (
	"context"
	"log"

	"github.com/node-real/bundle-go-sdk/client"
	"github.com/node-real/bundle-go-sdk/types"
)

func main() {
	privateKey := "<Your own private key>"
	account, err := types.NewAccountFromPrivateKey("test", privateKey)
	if err != nil {
		log.Fatalf("New account from private key error, %v", err)
	}

	rpcAddr := "https://gnfd-testnet-fullnode-tendermint-us.bnbchain.org:443"
	chainId := "greenfield_5600-1"
	
	gnfdCLient, err := client.New(chainId, rpcAddr, client.Option{DefaultAccount: account})
	if err != nil {
		log.Fatalf("unable to new greenfield bundleclient, %v", err)
	}
}

```

###  Quick Start Examples

The examples directory provides a wealth of examples to guide users in using the SDK's various features, including basic storage upload and download functions,
group functions, permission functions, as well as payment and cross-chain related functions.

The **basic.go** includes the basic functions to fetch the blockchain info.

The **storage.go** includes the most storage functions such as creating a bucket, uploading files, downloading files, heading and deleting resource.

The **group.go** includes the group related functions such as creating a group and updating group member.

The **payment.go** includes the payment related functions to manage payment accounts.

The **permission.go** includes the permission related functions to manage resources(bucket, object, group) policy.

The **crosschain.go** includes the cross chain related functions to transfer or mirror resource to BSC.


#### Config Examples

You need to modify the variables in "common.go" under the "examples" directory to set the initialization information for the client, including "rpcAddr", "chainId", and "privateKey", etc. In addition,
you also need to set basic parameters such as "bucket name" and "object name" to run the basic functionality of storage.

#### Run Examples
The steps to run example are as follows
```
make examples
cd examples
./storage 
```

You can also directly execute "go run" to run a specific example.
For example, execute "go run storage.go common.go" to run the relevant example for storage.
Please note that the "permission.go" example must be run after "storage.go" because resources such as objects need to be created first before setting permissions.

## Reference

- [Greenfield](https://github.com/bnb-chain/greenfield): the greenfield blockchain
- [Greenfield-Contract](https://github.com/bnb-chain/greenfield-contracts): the cross chain contract for Greenfield that deployed on BSC network.
- [Greenfield-Tendermint](https://github.com/bnb-chain/greenfield-tendermint): the consensus layer of Greenfield blockchain.
- [Greenfield-Storage-Provider](https://github.com/bnb-chain/greenfield-storage-provider): the storage service infrastructures provided by either organizations or individuals.
- [Greenfield-Relayer](https://github.com/bnb-chain/greenfield-relayer): the service that relay cross chain package to both chains.
- [Greenfield-Cmd](https://github.com/bnb-chain/greenfield-cmd): the most powerful command line to interact with Greenfield system.
- [Awesome Cosmos](https://github.com/cosmos/awesome-cosmos): Collection of Cosmos related resources which also fits Greenfield.