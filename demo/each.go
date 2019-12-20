package main

import (
	"io/ioutil"
	"fmt"
	"os"
)

func main() {
	i := []string{"F:/git/test/a","F:/git/test/b"}
	for {
		n := len(i)
		if n < 1 {
			break
		}
		dir := i[n-1]
		i = i[:n-1]
		// func ReadDir(dirname string) ([]os.FileInfo, error)
		f, err := ioutil.ReadDir(dir)
		if err != nil {
			panic(err)
		}
		for _, item := range f {
			file := fmt.Sprintf("%s/%s", dir, item.Name())
			info, err := os.Stat(file)
			if err != nil {
				panic(err)
			}
			if info.IsDir() {
				i = append(i, file)
				continue
			}
			fmt.Println(file)
		}
	}
	j := []string{"F:/git/test/a","F:/git/test/b"}
	println(len(j))
	fmt.Println(j[0:len(j)])

}