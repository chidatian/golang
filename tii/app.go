package tii

import(
	"fmt"
	"net/http"
)

type app struct {
	routes map[string]interface{}
}

func New() *app{
	return &app{
		make(map[string]interface{}),
	}
}

func (this *app) Router(pattern string, controller interface{}) {
	this.routes[pattern] = controller
}

func (this *app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(*r)
}

func (this *app) Run(port string) {
	fmt.Printf("Listen on %s\n", port)
	http.ListenAndServe(port, this)
}