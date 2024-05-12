package main

//"github.com/tendermint/tendermint/types"

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"

	blobstreamProof "github.com/CryptoKass/blobstreamx-example/blobstreamProof"
	"github.com/LinXJ1204/celestia-da-prover/VerifierAndTable"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

/* func main() {
	pf, _, err := blobstreamProof.GetSubmitInfoByTx("0FA4E6CE47A4337C09640B847947ECD40A306F5A6AD1DA264639AF55F634F360")
	if(err != nil) {
		return
	}
	verifyDataSubmitInContract(pf)

} */

// Alert!! using Celestia Mocha testnet and Arbitrum Sepolia testnet
const CelestiaMochaRpc = "https://celestia-mocha-rpc.publicnode.com:443"
const ArbitrumSepoliaRpc = "https://arbitrum-sepolia-rpc.publicnode.com"
const NewBlobstreamXContract = "0xc3e209eb245Fd59c8586777b499d6A665DF3ABD2"
const PayForBlobTxHash = "8C1F5B8372C8DB8E26CB5D76C2EA9F55094A32917DA5D80DD61F7E14247AF3B1"

// Warning: NewBlobstreamXContract updated every hour, so it's possible a data have be submitted to celestia but doesnt't update to NewBlobstreamXContract yet.
func verifyDataSubmitInContract(celestiaTx string) error {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
		return err
	}

	//Celestia transaction hash
	proof, _, err := blobstreamProof.GetSubmitInfoByTx(celestiaTx)
	fmt.Println(proof)
	client, err := ethclient.Dial(ArbitrumSepoliaRpc)
	if err != nil {
		return err
	}

	pk := os.Getenv("PK")
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		fmt.Println(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(nonce, gasPrice)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasPrice = gasPrice

	address := common.HexToAddress("0xA8974ABdf1F1E5E2256ba728a19cF6ffDb63af86")
	instance, err := VerifierAndTable.NewVerifierAndTable(address, client)
	if err != nil {
		fmt.Println(err)
	}
	tx, err := instance.VerifySharesToDataRootTupleRoot(auth, proof)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tx)

	return nil
}

func hashToByteArr(hash string) []byte {
	var hexStrings []string

	// Iterate over the string by steps of 2
	for i := 0; i < len(hash); i += 2 {
		// Check if there are at least two characters left
		if i+1 < len(hash) {
			// Add the substring of two characters to the result
			hexStrings = append(hexStrings, hash[i:i+2])
		} else {
			// If there is only one character left, add it with an empty string
			hexStrings = append(hexStrings, string(hash[i]))
		}
	}

	byteArray := make([]byte, len(hexStrings))
	for i, hexStr := range hexStrings {
		hexValue, err := hex.DecodeString(hexStr)
		if err != nil {
			fmt.Println("Error:", err)
			return []byte{}
		}
		byteArray[i] = hexValue[0]
	}

	return byteArray
}
