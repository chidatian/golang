package main

import (
	"fmt"
	"time"
)

func main() {
	// 格式化输出
	res := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(res)
	// 当前时间截
	res2 := time.Now().Unix()
	fmt.Println(res2)
	// 时间截 转 年月日 时分秒
	res3 := time.Unix(res2, 0).Format("2006-01-02 15:04:05")
	fmt.Println(res3)
}