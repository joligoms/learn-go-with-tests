package main

import (
	greeter "gowithtests/dependency-injection"
	"log"
	"net/http"
)

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	greeter.Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreetHandler)))
}
