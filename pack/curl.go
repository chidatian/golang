package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"flag"
)

type Curl struct {
	Status string
	Body []byte
	Num int
}

func (this *Curl) Get(url string) {
	res,err := http.Get(url)
	if err != nil {
		panic("something errors1")
	}
	this.Status = res.Status
	this.Body,_ = ioutil.ReadAll(res.Body)
}

func (this *Curl) Write(file string) {
	obj,err := os.Create(file)
	if err != nil {
		panic("something errors2")
	}
	this.Num,_ = obj.Write(this.Body)
}

var (
	url string
	file string
)

func init() {
	u := flag.String("u", "", "Input Your URL")
	f := flag.String("f", "", "Input Your FileName")
	flag.Parse()
	url,file = *u,*f
	if url == "" {
		panic("URL or FileName is  nil")
	}
}

func main() {
	c := new(Curl)
	c.Get(url)
	if file == "" {
		fmt.Println(string(c.Body))
	} else {
		c.Write(file)
		fmt.Println(c.Status)
	}
}