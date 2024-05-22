package main

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	jsoniter "github.com/json-iterator/go"

	"bundle-go-sdk/pkg/bundleclient"
)

func main() {
	endpointUrl := "https://bsc-testnet.nodereal.io/v1/" + os.Getenv("APIKEY")
	addressStr := os.Getenv("Address")
	privateKeyStr := os.Getenv("PrivateKey")

	address := common.Address(hexutil.MustDecode(addressStr))
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		panic(err)
	}

	rpcCli, err := rpc.Dial(endpointUrl)
	if err != nil {
		panic(err)
	}

	ethCli := ethclient.NewClient(rpcCli)
	bundleCli := bundleclient.New(rpcCli)

	// prepare bundle
	nonce, err := ethCli.NonceAt(
		context.Background(),
		address,
		big.NewInt(rpc.LatestBlockNumber.Int64()),
	)
	if err != nil {
		panic(err)
	}

	latestBlock, err := ethCli.BlockByNumber(context.Background(), big.NewInt(rpc.LatestBlockNumber.Int64()))
	if err != nil {
		panic(err)
	}

	fmt.Println("latest block number: ", latestBlock.Number(), "nonce: ", nonce)

	bundle := types.SendBundleArgs{
		Txs:               make([]hexutil.Bytes, 0),
		MaxBlockNumber:    0,
		MinTimestamp:      nil,
		MaxTimestamp:      nil,
		RevertingTxHashes: nil,
	}
	for i := 0; i < 3; i++ {
		txData := types.LegacyTx{
			Nonce:    nonce + uint64(i),
			To:       &address,
			Value:    big.NewInt(params.GWei),
			Gas:      uint64(5000000),
			GasPrice: big.NewInt(5e9),
			Data:     nil,
		}

		signer := types.MakeSigner(params.ChapelChainConfig, latestBlock.Number(), latestBlock.Time())

		tx, err := types.SignNewTx(privateKey, signer, &txData)
		if err != nil {
			panic(err)
		}

		rawTx, err := tx.MarshalBinary()
		if err != nil {
			panic(err)
		}
		bundle.Txs = append(bundle.Txs, rawTx)
	}

	// send bundle
	bundleHash, err := bundleCli.SendBundle(context.Background(), bundle)
	if err != nil {
		panic(err)
	}
	fmt.Println("bundle hash: ", bundleHash.String())

	// query bundle
	bundleQuery, err := bundleCli.QueryBundle(context.Background(), bundleHash)
	if err != nil {
		panic(err)
	}
	bundleJson, _ := jsoniter.Marshal(bundleQuery)
	fmt.Println("bundle queried: ", string(bundleJson))

	// bundle price
	bundlePrice, err := bundleCli.BundlePrice(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("bundle price: ", bundlePrice)

	// builders
	builders, err := bundleCli.Builders(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("builders: ", builders)

	// validators
	validators, err := bundleCli.Validators(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("validators: ", validators)
}
