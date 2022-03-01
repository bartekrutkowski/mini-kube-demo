package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	rand.Seed(time.Now().Unix())

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}

	return string(s)
}

func indexHandler(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("Hello world from %s\n", name)

		_, err := w.Write([]byte(msg))
		if err != nil {
			log.Fatal("error sending response to the request")
		}

		log.Println("sent response to incoming request")
	}
}

func main() {
	mux := http.NewServeMux()

	name := randomString(6)

	mux.HandleFunc("/", indexHandler(name))

	log.Printf("starting %s http server on port 8080\n", name)

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
