package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	log *log.Logger
}

func NewHello(log *log.Logger) *Hello {
	log.Println("builder")
	return &Hello{log}
}

func (hello *Hello) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	hello.log.Println("Hello World")

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Hello %s", data)
}
