package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Metadata struct {
	Description  string      `json:"description"`
	External_url string      `json:"external_url"`
	Image        string      `json:"image"`
	Name         string      `json:"name"`
	Attributes   []Attribute `json:"attributes"`
}

type Attribute struct {
	Trait_type string `json:"trait_type"`
	Value      string `json:"value"`
}

type Res struct {
	Id  int    `json:"id"`
	Res string `json:"res"`
}

const HOST = "https://psyduck-app.wayneies1206.workers.dev"

func main() {
	const file_count = 540
	for i := 0; i <= file_count; i++ {
		fmt.Println(i)
		content, err := ioutil.ReadFile(`/Users/zhengweilin/Downloads/outputsss/` + strconv.FormatUint(uint64(i), 10) + `.json`)
		if err != nil {
			log.Fatal("Error when opening file: ", err)
		}

		var payload Metadata
		err = json.Unmarshal(content, &payload)
		if err != nil {
			log.Fatal("Error during Unmarshal(): ", err)
		}

		if payload.Image != "" {
			fmt.Println("UPLOAD")
			uploadResult(i, payload.Image)
		}
	}
}

func uploadResult(id int, uri string) {
	jsonString := fmt.Sprintf("{\"URI\":\"%s\",\"id\":%d}", uri, id)
	_, err := http.Post(HOST+"/nftURIById", "application/json", bytes.NewReader([]byte(jsonString)))
	if err != nil {
		fmt.Println("Upload fail: %d", id)
	}
}

func getResult(id int) {
	jsonString := fmt.Sprintf("{\"id\":%d}", id)
	res, err := http.Post(HOST+"/getNftURIById", "application/json", bytes.NewReader([]byte(jsonString)))
	if err != nil {
		fmt.Println("Upload fail: %d", id)
	} else {
		fmt.Println(res.Body)

	}
}
