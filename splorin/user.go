package main 

import (
    "code.google.com/p/go.crypto/bcrypt"
    "database/sql"
    "fmt"
)

type User struct {
    id          int
    username    string
    email       string
    password    []byte
    salt        []byte
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

/*
*/
func Login(db *sql.DB, email, password string) (*User, error) {
    u, err := FindUser(db, email)
    if err != nil {
        fmt.Printf("Could not find user with that email: %s\n", email)
        return  nil, err
    }

    err = bcrypt.CompareHashAndPassword(u.password, []byte(password))
    if err != nil {
        fmt.Printf("Could not authenticate user\n")
        return nil, err
    }

    return u, err
}
