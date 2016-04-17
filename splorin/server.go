package main

import (
    "net/http"
    "fmt"
    "database/sql"
    "html/template"
    //"github.com/gorilla/securecookie"
    "github.com/gorilla/mux"
    //"log"
)

// Globals
var splorin_db *sql.DB

func main() {
    // Serve Mux
    r := mux.NewRouter()

    r.HandleFunc("/favicon.ico", favicon)
    r.HandleFunc("/login", login)
    r.HandleFunc("/authenticate", authenticate)
    r.HandleFunc("/register", register)
    r.HandleFunc("/createnewuser", createnewuser)
    r.HandleFunc("/", hello)

    http.Handle("/", r)

    splorin_db = Connect()
    defer Close(splorin_db)

    InitCookies()

    // Begin Serving
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        //log.Fatal("ListenAndServe: ", err)
        fmt.Printf("Couldn't serve!\n")
    }
}

func favicon(w http.ResponseWriter, r *http.Request) {
    // TODO: add something here
}

func login(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("html/login.html")
    if err != nil {
        fmt.Printf("void")
    }

    flash, err := GetFlash("auth", r, w)
    if (err != nil) {
        fmt.Printf("Auth Flash Error: %s\n", err)
    } else {
        fmt.Printf("Auth Flash: %s\n", flash)
    }

    err = t.Execute(w, nil)
}

func authenticate(w http.ResponseWriter, r *http.Request) {
    var redirect_url string
    flash_msg := ""

    user, err := AuthUser(splorin_db, w, r)

    if user == nil {
        flash_msg = "Username provided does not exist!"
    }

    if err != nil {
        redirect_url = "/login"
        fmt.Printf("Authentication Error: %s\n", err)
    } else {
        redirect_url = "/"
        SetCookieAuthenticated(user.email, w)

        fmt.Printf("Authentication Successful: [%s]\n", user.email)
    }

    SetFlash("auth", flash_msg, w)

    http.Redirect(w, r, redirect_url, http.StatusFound)
}

func register(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("html/register.html")
    if err != nil {
        fmt.Printf("void")
    }

    err = t.Execute(w, nil)
}

func createnewuser(w http.ResponseWriter, r *http.Request) {
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
    isAuthenticated := VerifyCookieAuthenticated(r)

    if isAuthenticated == false {
        fmt.Printf("Index Error: user is not authenticated!\n")
        //http.Redirect(w, r, "login", http.StatusFound)
        return
    }

    fmt.Printf("We found a match!\n")
    w.Write([]byte("hello!"))
}