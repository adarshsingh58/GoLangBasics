package main

import "net/http"

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello Go!!!"))
	})
	http.ListenAndServe(":8000", nil)

}
