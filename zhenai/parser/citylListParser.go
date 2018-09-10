package parser

import (
	"regexp"
	"007.com/爬虫/爬虫2/engines"
)

const regexpString = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>].+>([^<].+)</a>`

// 城市列表解析器
func CityListParser(b []byte) (r engines.RequestResult) {
	// ^[a-z]
	// [^[a-z]]
	// [a-z0-9]+
	//
	// class=""
	// >
	//[^<].+ 汉字城市名字
	match := regexp.MustCompile(regexpString)

	// 查找所有的匹配的字符串
	bytes := match.FindAllSubmatch(b, -1)

	for _, m := range bytes {
		//fmt.Printf("City:%s  URL:%s\n", m[2], m[1])
		//fmt.Println(r)
		r.Items = append(r.Items, m[2])
		r.R = append(r.R, engines.Request{
			string(m[1]),
			engines.NilParser,
		})
	}
	return
}
