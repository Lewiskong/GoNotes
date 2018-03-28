package main

import (
	"github.com/go-martini/martini"
	"net/http"
	"fmt"
)

func main() {
	m := martini.Classic()
	m.Post("/test", func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		fmt.Println(req.Form)
	})
	m.Run()
}
