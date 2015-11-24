// ローカルでHello word
package main

import (
	//"fmt"
	"net/http"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello, World"))
}

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}
