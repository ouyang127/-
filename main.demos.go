package main

import (
	"007.com/爬虫/爬虫2/zhenai/parser"
	"007.com/爬虫/爬虫2/engines"
)

func main() {
	//无限深入的去爬虫
	engines.Run(engines.Request{
		"http://www.zhenai.com/zhenghun",
		parser.CityListParser,
	})

}
