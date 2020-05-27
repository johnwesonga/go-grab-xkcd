package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/johnwesonga/go-grab-xkcd/client"
)

func main() {
	log.Println("Starting xkcd")
	httpListenPort := os.Getenv("PORT")
	if httpListenPort == "" {
		httpListenPort = "8080"
	}

	xkcdClient := client.NewXKCDClient()
	xkcdClient.SetTimeout(time.Duration(client.DefaultClientTimeout.Seconds()) * time.Second)

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		comic, err := xkcdClient.Fetch()
		if err != nil {
			log.Println(err)
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "<img src='%s'>", comic.Image)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", httpListenPort), nil))
}
