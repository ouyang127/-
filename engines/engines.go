package engines

import (
	"fmt"
	"log"
	"007.com/爬虫/爬虫2/fetcher"
)
//最后才走这里
func Run(r ...Request) {
	for len(r) > 0 {
		fristRequest := r[0]
		r = r[1:]
		fmt.Println("即将请求：", fristRequest.Url)
		all, err := fetcher.FetchData(fristRequest.Url)
		if err != nil {
			log.Printf("%v\n", err)
		}
		result := fristRequest.ParserFunc(all)
		r = append(r, result.R...)
	}
}
