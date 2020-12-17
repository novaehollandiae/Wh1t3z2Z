package main

import (
    "io"
    "log"
    "net/http"
)

func Tmp(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "version 1")
}

func main() {
    http.HandleFunc("/", Tmp)
    err := http.ListenAndServe(":3300", nil)
    if err != nil {
        log.Fatal(err)
    }
}
