package json

import (
	"encoding/json"
)

func Encode(obj interface{}) string{
	res,err := json.Marshal(obj)
	if err != nil {
		panic("something errors : json_encode")
	}
	return string(res)
}

func Decode(str string) interface{} {
	var r interface{}
	err := json.Unmarshal([]byte(str), &r)
	if err != nil {
		panic("something errors : json_decode")
	}
	return r
}