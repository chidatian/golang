package tii

import(
	"encoding/json"
)

func Encode(vlaue interface{}) string{
	s := json.MarShal(value)
	return s
}

func Decode(value string) interface{}{
	i := json.UnMarShal(value)
	return i
}