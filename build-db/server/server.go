package server

import (
	"log"
	"net/http"
)

func Start(host string, port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Helo mot"))
		log.Println("request: " + r.RequestURI)
	})
	err := http.ListenAndServe(host+":"+port, nil)
	if err != nil {
		log.Fatalln("listen error", err)
	}
}
