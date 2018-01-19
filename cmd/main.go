package main

import (
	"fmt"
	"imooc_crawler4go/common"
)

func main() {
	spider := common.ImoocSpider{URL: "https://coding.imooc.com/"}
	content, _ := spider.FetchContent()
	fmt.Println(content)
}
