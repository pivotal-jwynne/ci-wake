package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	startServer()
}

func startServer() {
	port := os.Getenv("PORT")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/stop", Stop)
	router.HandleFunc("/start", Start)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func Stop(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Stopping CI SAVER \n")
	tmp := []byte("")
	err := ioutil.WriteFile("../stay_awake", tmp, 0644)
	check(err)
}

func Start(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Starting CI SAVER \n")
	err := os.Remove("../stay_awake")
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
