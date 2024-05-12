package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const HOST = "https://psyduck-app.wayneies1206.workers.dev"

func UpdateNftOwnByAddress(receiver string, tokenId uint) {
	jsonString := fmt.Sprintf("{\"owner\": \"%s\", \"nftId\": %d, \"nftName\": \"none\", \"creatorAddr\": \"none\", \"url\": \"none\"}", receiver, tokenId)
	fmt.Println("UPDATE")
	res, err := http.Post(HOST+"/nftBuy", "application/json", bytes.NewReader([]byte(jsonString)))
	fmt.Println(res, err)
}

func UpdateBlock(blochHeight uint) {
	jsonString := fmt.Sprintf("{\"block\": %d}", blochHeight)
	res, err := http.Post(HOST+"/updateBlock", "application/json", bytes.NewReader([]byte(jsonString)))
	fmt.Println(res, err)
}

type BlockHeightResponse struct {
	LastupdateBlock string `json:"lastupdateBlock"`
}

func GetUpdatedBlock() int {
	var h BlockHeightResponse
	res, _ := http.Get(HOST + "/lastupdateBlock")
	err := json.NewDecoder(res.Body).Decode(&h)
	fmt.Println(err)
	num, err := strconv.Atoi(h.LastupdateBlock)
	fmt.Println(num)
	return num
}
