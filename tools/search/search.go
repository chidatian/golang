package search

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"regexp"
)

type Result struct {
	Field string
	Content string
}

var (
	nulls []string
	notxml []string
	errors []error
)

var result = make(chan *Result)

func Run(sear string) {
	src := Feed()
	var pool sync.WaitGroup
	go func() {
		for {
			select {
			case item := <-result :
				fmt.Println(item)
			}
		}
	}()
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
			if xml.hasErr {
				notxml = append(notxml, v.Link)
				return
			}
			Match(sear, xml, v)
		}(v)
	}
	pool.Wait()
	fmt.Println(nulls)
	fmt.Println(notxml)
	fmt.Println(errors)
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
	var (
		isB bool
		err error
	)
	for _, item := range xml.Chann.Items {
		isB, err = regexp.MatchString(sear, item.Title)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		if isB {
			result <- &Result{
				Field: "Title",
				Content: item.Title,
			}
			/*fmt.Println(Result{
				Field: "Title",
				Content: item.Title,
			})*/
		}
		isB, err = regexp.MatchString(sear, item.Description)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		if isB {
			result <- &Result{
				Field: "Description",
				Content: item.Description,
			}
			/*fmt.Println(Result{
				Field: "Description",
				Content: item.Description,
			})*/
		}
	}
}
