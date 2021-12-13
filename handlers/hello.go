package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")
	log.Println("hello world")
	d, err := ioutil.ReadAll(r.Body) //ioreadcloser
	if err != nil {
		// rw.WriteHeader(http.StatusBadRequest)
		// rw.Write([]byte("Oooopps"))
		http.Error(rw, "ooops", http.StatusBadRequest)
		return
	}
	log.Printf("Data %s\n", d)
	fmt.Fprintf(rw, "Hello %s", d)
}
