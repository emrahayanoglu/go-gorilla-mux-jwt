package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func LoginHandler() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if token, err := CreateToken(); err != nil {
			fmt.Fprintf(w, "Error in Token Creation")
		} else {
			testmap := make(map[string]string)
			testmap["token"] = token
			b, err := json.Marshal(testmap)
			if err != nil {
				fmt.Fprintf(w, "Error in JSON Marshalling")
			}

			fmt.Fprintf(w, string(b))
		}
	}

	return http.HandlerFunc(fn)
}

func HelloHandler() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		testmap := make(map[string]string)
		testmap["hello"] = "world in authorized page"
		b, err := json.Marshal(testmap)
		if err != nil {
			fmt.Fprintf(w, "Error in JSON Marshalling")
		}

		fmt.Fprintf(w, string(b))
	}

	return http.HandlerFunc(fn)
}

func RunServer(port int, done chan bool) {
	r := mux.NewRouter()

	r.HandleFunc("/login", LoginHandler())

	memberRouter := r.PathPrefix("/member").Subrouter().StrictSlash(true)
	memberRouter.Use(AuthenticationMiddleware)
	memberRouter.HandleFunc("/hello", HelloHandler())

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
