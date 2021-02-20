package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := "hello world"
	if _, err := w.Write([]byte(message)); err != nil {
		log.Println(err)
	}
}

func traceLog(w http.ResponseWriter, r *http.Request) {
	if body, err := ioutil.ReadAll(r.Body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		log.Println(string(body))
	}
}

func trace(w http.ResponseWriter, r *http.Request) {
	if payload, err := httputil.DumpRequest(r, true); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		log.Println(string(payload))
	}
}

func main() {
	http.HandleFunc("/sayhello", sayHello)
	http.HandleFunc("/log", traceLog)
	http.HandleFunc("/log1", trace)
	if err := http.ListenAndServe("0.0.0.0:3000", nil); err != nil {
		log.Fatalln(err)
	}
}
