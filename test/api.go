package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	resp, err := http.PostForm("http://www.zzybgp.com/api/user/do_main/select",url.Values{"key": {"Value"}, "id": {"123"}})
	if err != nil {
		panic("something errors: post")
	}

	fmt.Println(resp)
}