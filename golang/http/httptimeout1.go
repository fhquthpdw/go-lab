package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	db := tt{
		"foo": "football",
		"bar": "basketball",
	}
	log.Println("listing port: 3030")

	s := http.Server{
		Addr:         ":3030",
		WriteTimeout: 1 * time.Second,
		//ReadTimeout:  10 * time.Second,
		Handler: http.TimeoutHandler(db, 5*time.Second, "time out"),
	}

	log.Fatal(s.ListenAndServe())
}

type tt map[string]string

func (db tt) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
	fmt.Fprintf(w, "%s, %s", r.Method, r.RequestURI)
	for k, v := range db {
		fmt.Fprintf(w, "%s: %s", k, v)
	}
}
