package main

import "net/http"

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":80", func(h http.Handler) http.Handler {
		println("Server is running on port 80")
		println("Try: curl http://localhost")
		println("Press Ctrl-C to stop")
		return h
	}(http.DefaultServeMux))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
