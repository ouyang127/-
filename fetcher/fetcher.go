package fetcher//倒数第二步

import (
	"net/http"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
	"io/ioutil"
	"io"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/text/encoding/unicode"
)

func FetchData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error:Status Code:%d", resp.StatusCode)
	}
	e := determineEncoding(resp.Body)

	utf8Reader := transform.NewReader(resp.Body,
		e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
