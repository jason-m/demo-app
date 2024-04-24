package main

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
)

func main() {
	listenIP := net.ParseIP("127.0.0.1")
	listenPort := 8080

	listenString := listenIP.String() + ":" + strconv.Itoa(listenPort)

	http.HandleFunc("/", webEcho)
	http.ListenAndServe(listenString, nil)

}

func webEcho(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	w.WriteHeader(400)
	for key, values := range r.Form {
		for _, value := range values {
			fmt.Fprintf(w, "GET Form - Key: %s, Value: %s\n", key, value)
		}
	}

	for key, values := range r.PostForm {
		for _, value := range values {
			fmt.Fprintf(w, "POST Form - Key: %s, Value: %s\n", key, value)
		}
	}

}
