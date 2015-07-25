package main

import (
    "net/http"
    "fmt"
    "database/sql"
    "html/template"
    "github.com/gorilla/sessions"
    //"log"
)

// Globals
var splorin_db *sql.DB
var store = sessions.NewCookieStore([]byte(SESSION_PASSWORD))

func main() {

    // Serve Mux
    http.HandleFunc("/", hello)
    http.HandleFunc("/login", login)
    http.HandleFunc("/authenticate", authenticate)

    splorin_db = Connect()
    defer Close(splorin_db)

    user := CreateNewUser("what@yo.com", "lolpass")
    err := AddUser(splorin_db, user)
    if err != nil {
        fmt.Printf("Error adding a new user!\n")
    }

    // Begin Serving
    err = http.ListenAndServe(":8080", nil)
    if err != nil {
        //log.Fatal("ListenAndServe: ", err)
        fmt.Printf("Couldn't serve!\n")
    }


}

func login(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("html/login.html")
    if err != nil {
        fmt.Printf("void")
    }

    err = t.Execute(w, nil)

    //user, err := AuthUser(splorin_db, w, r)
    //_ = user
}

func authenticate(w http.ResponseWriter, r *http.Request) {
    user, err := AuthUser(splorin_db, w, r)
    if err != nil {
        fmt.Printf("Authenticate Error: %s\n", err)
    }
    _ = user
    _ = err
}

func register(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("void")
}

func hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello!"))
}