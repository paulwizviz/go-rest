package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {

	r := chi.NewRouter()

	r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Root hello"))
	})

	r.Route("/v1", func(r chi.Router) {

		r.Use(func(next http.Handler) http.Handler {
			fn := func(rw http.ResponseWriter, r *http.Request) {
				log.Println("First middleware call")
				next.ServeHTTP(rw, r)
			}
			return http.HandlerFunc(fn)
		})

		r.Use(func(next http.Handler) http.Handler {
			fn := func(rw http.ResponseWriter, r *http.Request) {
				log.Println("Second middleware call")
				next.ServeHTTP(rw, r)
			}
			return http.HandlerFunc(fn)
		})

		r.Get("/hello", func(rw http.ResponseWriter, r *http.Request) {
			rw.Write([]byte("New Hello"))
		})

		r.Get("/ola", func(rw http.ResponseWriter, r *http.Request) {
			rw.Write([]byte("Ola"))
		})
	})

	log.Println("Starting on port: 3000")
	http.ListenAndServe("0.0.0.0:3000", r)
}
