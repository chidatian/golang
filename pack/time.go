<<<<<<< HEAD
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
=======
package main

import (
	"fmt"
	"time"
)

func main() {
	// 格式化输出
	// func Now() Time
	// func (t Time) Format(layout string) string
	res := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(res)
	// 当前时间截
	// func (t Time) Unix() int64
	res2 := time.Now().Unix()
	fmt.Println(res2)
	// 时间截 转 年月日 时分秒
	res3 := time.Unix(res2, 0).Format("2006-01-02 15:04:05")
	fmt.Println(res3)
>>>>>>> 65888be0467c07f28b57cd4073746ad38fcc0782
}