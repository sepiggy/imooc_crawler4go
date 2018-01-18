package common

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

const URL string = "https://coding.imooc.com"
