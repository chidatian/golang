package pack

// 堆栈 test
type Stacker interface {
	Pop() string		// 出栈
	Push(val string)	// 入栈
	Clear()
}

type Stack struct {
	Top *Child
	Count int
}

type Child struct {
	Value string
	Next *Child
}

func New() *Stack{
	return new(Stack)
}

func (this *Stack) Pop() string {
	res := this.Top.Value
	this.Top = this.Top.Next
	this.Count--
	return res
}

func (this *Stack) Push(val string) {
	c := Child{}
	c.Value = val
	if this.Top == nil {
		this.Top = &c
	} else {
		c.Next = this.Top
		this.Top = &c
	}
	this.Count++
}

func (this *Stack) Clear() {
	this.Top = nil
}