package main

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/LinXJ1204/celestia-da-prover/drawableNft721"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	ctx := context.Background()

	// step 0. get eth-rpc and trpc endpoint
	ethEndpoint := "wss://arbitrum-sepolia.infura.io/ws/v3/3db317632844470fa86a5a3a8cb7724d"
	contractAddress := "0x585a1DDaB9116F483d367bCa6eb64797252051c8"

	eth, err := ethclient.Dial(ethEndpoint)
	if err != nil {
		fmt.Println(err)
	}

	contract, err := drawableNft721.NewDrawableNft721(common.HexToAddress(contractAddress), eth)
	if err != nil {
		fmt.Println(err)
	}

	for {

		start := uint64(GetUpdatedBlock()) + 1
		if start == 0 {
			return
		}

		bh := start
		fmt.Println(start)

		pp, err := contract.FilterTransfer(&bind.FilterOpts{
			Start:   start,
			Context: ctx,
		}, []common.Address{}, []common.Address{}, []*big.Int{})
		if err != nil {
			fmt.Println(err)
			return
		}

		for pp.Next() {
			event := pp.Event
			// Print out all caller addresses
			fmt.Printf(event.TokenId.String())
			bh = event.Raw.BlockNumber
			UpdateNftOwnByAddress(event.To.Hex(), uint(event.TokenId.Int64()))
			fmt.Println("\n")
		}

		UpdateBlock(uint(bh))

		time.Sleep(time.Duration(5) * time.Second)
	}

}
