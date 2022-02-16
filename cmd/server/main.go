package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	http.HandleFunc("/", gate)

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func gate(_ http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))
}
