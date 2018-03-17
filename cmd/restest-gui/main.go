package main

import "net/http"

func main() {
	// static server
	staticS := &http.Server{
		Addr:    ":8080",
		Handler: http.FileServer(http.Dir("/Users/lfree/Works/github.com/src/github.com/AYJiaYou/restest-gui/web/dist")),
	}
	staticS.ListenAndServe()
}
