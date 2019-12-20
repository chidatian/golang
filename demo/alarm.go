package main

import(
	"fmt"
	"os"
	"time"
)

const FILENAME = "./alarm.log"

func create() *os.File {
	// func Create(name string) (file *File, err error)
	f, err := os.Create(FILENAME)
	if err != nil {
		panic(err)
	}
	return f
}

func isExist() bool {
	_, err := os.Stat(FILENAME)
    return err == nil || os.IsExist(err)
}

func open() *os.File {
	// func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
	f,err := os.OpenFile(FILENAME, os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}
	return f
}

func append(f *os.File) {
	// func (f *File) WriteString(s string) (ret int, err error)
	t := time.Now().Format("2006-01-02 15:04:05")
	_,err := f.WriteString(fmt.Sprintf("%s\n", t))
	if err != nil {
		panic(err)
	}
}

func main() {
	var f *os.File
	if isExist() {
		f = open()
	} else {
		f = create()
	}
	append(f)
}