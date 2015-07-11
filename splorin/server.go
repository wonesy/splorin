package main

import (
    "net/http"
    "splorin/db"
    //"log"
)

func main() {

    // Serve Mux
    http.HandleFunc("/", hello)

    // Begin Serving
    /*
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
    */
    db.AddUser("Cameron")
}

func hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello!"))
}