package message

import (
	"bufio"
	"os"
)

type Msg struct {}

func (this *Msg) InitBuf() string{
	var input *bufio.Reader
	input = bufio.NewReader(os.Stdin)
	content, err := input.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return content
}