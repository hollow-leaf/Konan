package blobstreamProof

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/CryptoKass/blobstreamx-example/client"
	shareloader "github.com/CryptoKass/blobstreamx-example/contracts/ShareLoader.sol"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tendermint/tendermint/rpc/client/http"
)

func GetSubmitInfoByTx(_txHash string) (*shareloader.SharesProof, [32]byte, error) {
	ctx := context.Background()

	// step 0. get eth-rpc and trpc endpoint
	ethEndpoint := "https://arbitrum-sepolia-rpc.publicnode.com"
	trpcEndpoint := "https://celestia-mocha-rpc.publicnode.com:443"

	// See contract here: https://sepolia.arbiscan.io/contract/0xf2787995D9eb43b57eAdB361227Ddf4FEC99b5Df
	// This is the address of the ShareLoader contract
	// The contract wraps the DAVerifier.verifySharesToDataRootTupleRoot method
	//shareloaderAddress := common.HexToAddress("0xf2787995D9eb43b57eAdB361227Ddf4FEC99b5Df")

	// step 1: connect to eth and trpc endpoints
	eth, err := ethclient.Dial(ethEndpoint)
	if err != nil {
		return nil, [32]byte{}, err
	}
	trpc, err := http.New(trpcEndpoint, "/websocket")
	if err != nil {
		return nil, [32]byte{}, err
	}
	fmt.Println("Successfully connected to the Ethereum node and Tendermint RPC")

	// step 2: generate share proof
	txHash, err := hex.DecodeString(_txHash)
	tx, err := trpc.Tx(ctx, txHash, true)
	if err != nil {
		fmt.Println(err)
		return nil, [32]byte{}, err
	}

	fmt.Println(tx.Proof.ShareProofs[0].Start)

	sp := &client.SharePointer{
		Height: tx.Height,
		Start:  int64(tx.Proof.ShareProofs[0].Start),
		End:    int64(tx.Proof.ShareProofs[0].End),
	}

	//Using this proof to contract verify
	proof, root, err := client.GetShareProof(eth, trpc, sp)
	if err != nil {
		return nil, [32]byte{}, err
	}
	fmt.Println("Successfully generated share proof")

	return proof, root, nil

	// step 3: verify the share proof
	/* loader, err := shareloader.NewShareloader(shareloaderAddress, eth)
	if err != nil {
		panic(fmt.Errorf("failed to instantiate shareloader contract: %w", err))
	}

	valid, errCodes, err := loader.VerifyShares(nil, *proof, root)
	if err != nil {
		panic(fmt.Errorf("failed to verify share proof: %w", err))
	}

	// step 4: print the result
	if !valid {
		fmt.Println("Proof is invalid", "Error codes:", errCodes)
	}
	fmt.Println("Proof is valid") */
}
