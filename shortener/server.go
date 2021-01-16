package main

import (
	"fmt"
	"html"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {

	http.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Redirect, %q", html.EscapeString(r.URL.Path))
	})

	s := &http.Server{
		Addr:           ":5050",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	logrus.Info("Server started on: ", s.Addr)
	logrus.Fatal(s.ListenAndServe())
}
