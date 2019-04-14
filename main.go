package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prem0132/spark-operator-livy/pkg/v1alpha1"
	//"gopkg.in/yaml.v2"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/echo", v1alpha1.EchoHandler)
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/todos", TodoIndex)
	mux.HandleFunc("/todos/{todoId}", TodoShow)
	mux.HandleFunc("/batches", v1alpha1.Batches)
	mux.HandleFunc("/hello", v1alpha1.Hello)

	log.Fatal(http.ListenAndServe(":8998", mux))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Todo Index!")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
