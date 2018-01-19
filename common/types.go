package common

import (
	"net/http"
	"io/ioutil"
	"log"
)

// 课程
type Course struct {
	Name   string
	Grade  string
	Number int
	Desc   string
	Price  float32
}

// 一级类别
type FirstCategory struct {
	Name string
	Abbr string
}

// 二级类别
type SecondCategory struct {
	FirstCategory
	Name string
	Abbr string
}

// 爬虫接口
type Spider interface {
	FetchContent() (string, error)
	Analysis(string) string
	Refine(string) string
	Sort(string) string
	Show(string)
}

// 慕课爬虫结构
type ImoocSpider struct {
	URL           string
	RootPattern   string
	NamePattern   string
	NumberPattern string
}

func (self ImoocSpider) FetchContent() (string, error) {
	resp, err := http.Get(self.URL)
	if err != nil {
		log.Fatalf("FetchContent Error: %v", err)
		return "", err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("FetchContent Error: %v", err)
		return "", err
	}
	return string(data), err
}

func (self ImoocSpider) Analysis(htmls string) string {
	return ""
}

func (self ImoocSpider) Refine(htmls string) string {
	return ""
}

func (self ImoocSpider) Sort(htmls string) string {
	return ""
}

func (self ImoocSpider) Show(anchors string) string {
	return ""
}

