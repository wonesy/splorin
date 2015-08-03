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
    session, err := store.Get(r, SESSION_PASSWORD)
    if err != nil {
        return nil, err
    }

    err = r.ParseForm()
    if err != nil {
        return nil, err
    }

    email := r.PostFormValue("email")
    password := r.PostFormValue("password")

    fmt.Printf(email)
    
    user, err := FindUser(db, email)
    if user != nil {
        session.AddFlash("That email is already registered!")
        session.Save(r, w)
        return nil, err
    }

    user = CreateNewUser(email, password)
    err = AddUser(db, user)
    if err != nil {
        session.AddFlash("Could not add user!")
        session.Save(r, w)
        return nil, err
    }

    return user, err
}

/*
*/
func AuthUser(db *sql.DB, w http.ResponseWriter, r *http.Request) (*User, error) {
    session, err := store.Get(r, SESSION_PASSWORD)
    if err != nil {
        return nil, err
    }

    err = r.ParseForm()
    if err != nil {
        return nil, err
    }

    email := r.PostFormValue("email")
    password := r.PostFormValue("password")

    u, err := FindUser(db, email)
    if err != nil {
        session.AddFlash("That email has not been registered.\n")
        session.Save(r, w)
        return  nil, err
    }

    err = bcrypt.CompareHashAndPassword(u.password, []byte(password))
    if err != nil {
        session.AddFlash("Incorrect credentials.\n")
        session.Save(r, w)
        return nil, err
    }

    session.Values["userID"] = u.id
    session.Save(r, w)

    // TODO redirect user to page they came from

    return u, err
}
