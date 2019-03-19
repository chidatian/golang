package tii

type router struct {
	routes map[string]interface{}
}

func (this *router) Router(pattern string, controller interface{}) {
	this.routes[pattern] = controller
}