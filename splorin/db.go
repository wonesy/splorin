package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

/*
    Postgres Setup

    CREATE DATABASE splorin;

    CREATE TABLE users (
        id serial primary key,
        email varchar(100) not null unique,
        password varchar(128)
    );
*/

/*
    Function: Connect

    Establishes connection to the database

    Returns:

        Pointer to the connected database
*/
func Connect() (db *sql.DB) {
    dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
        DB_USER, DB_PASSWORD, DB_NAME)

    db, err := sql.Open("postgres", dbinfo)
    if err != nil {
        panic(err)
    }

    return db
}

/*
    Function: Close

    Closes a database connection

    Parameters:

        db  - Pointer to an open database connection

    Returns:

        Nothing
*/
func Close(db *sql.DB) {
    db.Close()
}

/*
    Function: AddUser

    Adds a user (and associated information) to the database

    Parameters:

        db      - Pointer to an open database instance
        name    - User's username
        email   - User's email

    Returns:

        Database error
*/
func AddUser(db *sql.DB, user *User) error {
    fmt.Printf("Adding user:\n\tEmail: %s\n", user.email)

    var lastInsertID uint64
    err := db.QueryRow("INSERT INTO users(email,password) VALUES($1,$2) returning id;",
        user.email, user.password).Scan(&lastInsertID)

    if err != nil {
        fmt.Printf("Error: %s\n", err)
    }

    return err
}


/*
    Function: FindUser

    Searches for a particular user in the database by email

    Parameters:

        db      - Pointer to an open database instance
        email   - User's email

    Returns:

        Pointer to User instance
        Error code
*/
func FindUser(db *sql.DB, email string) (*User, error) {
    var u User

    err := db.QueryRow("SELECT * FROM users WHERE email=$1;", 
        email).Scan(&u.id, &u.email, &u.password)
    if err != nil {
        fmt.Println(err)
        return &u, err
    }

    return &u, err
}
