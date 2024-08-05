package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	server := http.Server{
		Addr: ":8081",
	}

	http.HandleFunc("/execute", Handler)

	server.ListenAndServe()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 300*time.Millisecond)
	defer cancel()

	req, _ := http.NewRequest("GET", "http://localhost:8080/execute", nil)
	req = req.WithContext(ctx)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("err: " + err.Error())
		return
	}

	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)
	w.Write([]byte(respBody))
}
