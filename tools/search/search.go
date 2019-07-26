package search

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var (
	nulls []string

)

func Run(sear string) {
	src := Feed()
	var pool sync.WaitGroup
	for _,v := range src {
		pool.Add(1)
		go func(v Source) {
			defer pool.Done()
			res := Curl(v.Link)
			if res == "" {
				nulls = append(nulls, v.Link)
				return
			}
			xml := NewRss(res)
			Match(sear, xml, v)
		}(v)
		break
	}
	pool.Wait()
	fmt.Println(nulls)
}

func Curl(link string) string {
	resp, err := http.Get(link)
	if err != nil {
		return ""
	}
	re, er := ioutil.ReadAll(resp.Body)
	if er != nil {
		return ""
	}
	resp.Body.Close()
	return string(re)
}

func Match(sear string, xml Rss, src Source) {
	for k, v := range xml.Chann.Items {
		fmt.Println(k, v)
	}
	fmt.Println(src)
	fmt.Println(sear)
}
