package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func TestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		fmt.Println("Test Middleware is executed!")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func TestMiddleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		fmt.Println("Test Middleware 2 is executed!")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func HelloHandler() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		testmap := make(map[string]string)
		testmap["hello"] = "world"
		b, err := json.Marshal(testmap)
		if err != nil {
			fmt.Fprintf(w, "Error in JSON Marshalling")
		}

		fmt.Fprintf(w, string(b))
	}

	return http.HandlerFunc(fn)
}

func HelloDirect() func(http.ResponseWriter, *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(r.Header)

		testmap := make(map[string]string)
		testmap["hello"] = "world"
		b, err := json.Marshal(testmap)
		if err != nil {
			fmt.Fprintf(w, "Error in JSON Marshalling")
		}

		fmt.Fprintf(w, string(b))
	}

	return fn
}

func RunServer(port int, done chan bool) {
	r := mux.NewRouter()

	r.HandleFunc("/hello", HelloHandler())
	r.Handle("/hello2", TestMiddleware(HelloHandler()))

	r.HandleFunc("/hello3", HelloDirect())

	subrouter := r.PathPrefix("/sub/").Subrouter()

	subrouter.Use(TestMiddleware)
	subrouter.Use(TestMiddleware2)
	subrouter.HandleFunc("/hello", HelloHandler())

	err := http.ListenAndServe(":"+strconv.Itoa(port), r)

	if err != nil {
		panic(err)
	}

	done <- true
}

func main() {
	done := make(chan bool)

	RunServer(8080, done)

	<-done
}
