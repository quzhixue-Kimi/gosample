package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	if res, err := http.Get("http://localhost:3000/sayhello"); err != nil {
		fmt.Println(err)
		return
	} else {
		defer res.Body.Close()
		if n, err := ioutil.ReadAll(res.Body); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(n))
		}
	}
}
