package tcc

import (
	"fmt"
	"net/http"
	"reflect"
)

type App struct {
	Routes map[string]interface{}
}

func New() *App {
	return &App{
		Routes: make(map[string]interface{}),
	}
}

func (this *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	/*fmt.Println(r.URL.Query())
	fmt.Println(r.Proto)
	fmt.Println(r.RemoteAddr)
	fmt.Println(r.RequestURI)
	fmt.Println(r.Method)
	res,_ := ioutil.ReadAll(r.Body)
	fmt.Println(string(res))*/
	get := r.URL.Query()
	controller := get.Get("controller")
	action := get.Get("action")
	if this.Routes[controller] == nil {
		fmt.Println("something errors : controller by param")
	} else {
		if action == "" {
			action = "Index"
		}
		c := this.Routes[controller]
		e := reflect.ValueOf(c).Elem()
		e.FieldByName("Request").Set(reflect.ValueOf(r))
		e.FieldByName("Response").Set(reflect.ValueOf(w))
		res := e.MethodByName(action)
		if res.IsValid() {
			res.Call([]reflect.Value{})
		}
	}
}

func (this *App) Run(port string) {
	fmt.Println("listen : ", port)
	fmt.Println(this.Routes)
	var s http.Server
	s.Addr = port
	s.Handler = this
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}