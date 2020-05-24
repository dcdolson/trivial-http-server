package main

import (
    "log"
    "net/http"
    "os"
)

func main() {
        if len(os.Args) < 2 {
            log.Fatal("Usage: %v <folder>", os.Args[0])
        }
        folder := os.Args[1]
        fs := http.FileServer(http.Dir(folder))
        http.Handle("/", fs)
	log.Printf("About to listen on 8000. Go to https://127.0.0.1:8000/")
	err := http.ListenAndServe(":8000", nil)
	log.Fatal(err)
}

