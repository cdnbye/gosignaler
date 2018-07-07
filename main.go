package main

import (
	"flag"
	"net/http"
	"log"
)

var addr = flag.String("addr", ":8081", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {

	flag.Parse()
	hub := newHub()
	go hub.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Printf("URL: %s\n", r.URL.String())
		r.ParseForm()
		id := r.Form.Get("id")
		if id != "" {
			serveWs(hub, w, r, id)
		}
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}


}
