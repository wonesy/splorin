package main 

import (
    "code.google.com/p/go.crypto/bcrypt"
    "database/sql"
    "net/http"
    "fmt"
)

type User struct {
    id          int
    email       string
    password    []byte
}

/*
*/
func CreateNewUser(email, password string) (*User) {
    var user User

    fmt.Printf("Creating new user.\n")

    user.email = email
    user.SetPassword(password)

    return &user
}

/*
*/
func (u *User) SetPassword(password string) {
    hpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }

    u.password = hpass
}

func RegisterUser(db *sql.DB, w http.ResponseWriter, r *http.Request) (*User, error) {
    err := r.ParseForm()
    if err != nil {
        return nil, err
    }

    email := r.PostFormValue("email")
    password := r.PostFormValue("password")

    fmt.Printf(email)
    
    user, err := FindUser(db, email)
    if user != nil {
        return nil, err
    }

    user = CreateNewUser(email, password)
    err = AddUser(db, user)
    if err != nil {
        return nil, err
    }

    return user, err
}

/*
*/
func AuthUser(db *sql.DB, w http.ResponseWriter, r *http.Request) (*User, error) {
    err := r.ParseForm()
    if err != nil {
        return nil, err
    }

    email := r.PostFormValue("email")
    password := r.PostFormValue("password")

    u, err := FindUser(db, email)
    if err != nil {
        fmt.Printf("ERR: Could not find user")
        return  nil, err
    }

    err = bcrypt.CompareHashAndPassword(u.password, []byte(password))
    if err != nil {
        fmt.Printf("ERR: passwords did not match")
        return nil, err
    }

    return u, err
}
