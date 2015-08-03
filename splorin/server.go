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
    http.HandleFunc("/register", register)
    http.HandleFunc("/createnewuser", createnewuser)

    splorin_db = Connect()
    defer Close(splorin_db)

    // Begin Serving
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        //log.Fatal("ListenAndServe: ", err)
        fmt.Printf("Couldn't serve!\n")
    }


}

func login(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, SESSION_PASSWORD)
    
    flashes := session.Flashes()
    if len(flashes) > 0 {
        fmt.Printf("%v", flashes)
    }

    t, err := template.ParseFiles("html/login.html")
    if err != nil {
        fmt.Printf("void")
    }

    err = t.Execute(w, nil)
}

func authenticate(w http.ResponseWriter, r *http.Request) {
    user, err := AuthUser(splorin_db, w, r)
    if err != nil {
        fmt.Printf("Authenticate Error: %s\n", err)
    }
    
    _ = user

    var redirect_url string
    if err != nil {
        redirect_url = "login"
    } else {
        redirect_url = "/"
    }

    http.Redirect(w, r, redirect_url, http.StatusFound)
}

func register(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, SESSION_PASSWORD)

    flashes := session.Flashes()
    if len(flashes) > 0 {
        fmt.Printf("%v", flashes)
    }

    t, err := template.ParseFiles("html/register.html")
    if err != nil {
        fmt.Printf("void")
    }

    err = t.Execute(w, nil)
}

func createnewuser(w http.ResponseWriter, r*http.Request) {
    user, err := RegisterUser(splorin_db, w, r)
    if err != nil {
        fmt.Printf("Registration Error: %s\n", err)
    }
    
    _ = user

    var redirect_url string
    if err != nil {
        redirect_url = "register"
    } else {
        redirect_url = "/"
    }

    http.Redirect(w, r, redirect_url, http.StatusFound)
}

func hello(w http.ResponseWriter, r *http.Request) {
    session, err := store.Get(r, SESSION_PASSWORD)

    fmt.Printf(session.Values["userID"])
    w.Write([]byte("hello!"))
}