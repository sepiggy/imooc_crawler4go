package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	"imooc_crawler4go/common"
)

func main() {
	resp, err := http.Get(common.URL)
	checkError(err)
	data, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	fmt.Printf("Got:%q", string(data))
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Get : %v", err)
	}
}
