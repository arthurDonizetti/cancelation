package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {

	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/execute", Handler)

	server.ListenAndServe()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	ch := make(chan string)

	go func() {
		i := 1
		for {
			select {
			case <-ctx.Done():
				log.Printf("request canceled")
				close(ch)
				return
			default:
				log.Print(i)
				if i == 10 {
					ch <- "Process done"
					return
				}
				i++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	select {
	case res := <-ch:
		log.Printf("log " + res)
		w.Write([]byte(res))
	case <-ctx.Done():
		log.Printf("context closed")
		w.WriteHeader(499)
	}
}
