package main

import (
	"fmt"
	"tools/json"
)

var (
	j string = `[{"id":1,"name":"zouchunlei","pwd":123},{"id":2,"name":"qwe","pwd":123},{"id":3,"name":"qwe","pwd":123},{"id":4,"name":"qwe","pwd":123},{"id":5,"name":1,"pwd":null},{"id":6,"name":"qw22","pwd":123},{"id":7,"name":"qwe222","pwd":121},{"id":8,"name":"test","pwd":123},{"id":9,"name":"this is test","pwd":123},{"id":10,"name":222222222,"pwd":123},{"id":11,"name":"lixiao","pwd":123},{"id":12,"name":"qwertyu","pwd":123},{"id":13,"name":"xxx","pwd":"xxx"},{"id":14,"name":"oo","pwd":123},{"id":15,"name":"test10","pwd":123},{"id":16,"name":"test47","pwd":123},{"id":17,"name":"aaa","pwd":"bbb"},{"id":18,"name":"bbb","pwd":"aaa"},{"id":19,"name":"test711","pwd":123456},{"id":20,"name":"test700","pwd":123456}]`
)

func main() {
	res := json.Decode(j)
	fmt.Println(res)
	re := res.([]interface{})
	for k,v := range re {
		val := v.(map[string]interface{})
		fmt.Println(k,val["name"])
	}
}