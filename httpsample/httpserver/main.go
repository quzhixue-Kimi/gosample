package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := "hello world"
	if _, err := w.Write([]byte(message)); err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/sayhello", sayHello)
	if err := http.ListenAndServe("127.0.0.1:3000", nil); err != nil {
		fmt.Println(err)
	}
}
