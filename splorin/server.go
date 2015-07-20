package main

import (
    "net/http"
    "fmt"
    "database/sql"
    "github.com/gorilla/sessions"
    //"log"
)

// Globals
var splorin_db *sql.DB
var store = sessions.NewCookieStore([]byte(SESSION_PASSWORD))

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
    splorin_db = Connect()
    defer Close(splorin_db)

    err := AddUser(splorin_db, "cameron", "what@yo.com")
    if err != nil {
        fmt.Printf("Error adding a new user!\n")
    }

    //usr := &User{}
    usr, err := Login(splorin_db, "what@yo.com", "passss")
    _ = usr

}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("void")
}

func register(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("void")
}

func hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello!"))
}