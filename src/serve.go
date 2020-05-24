package main

import (
    "log"
    "net/http"
    "os"
)

func main() {
        if len(os.Args) < 3 {
            log.Fatal("Usage: %v <folder> port=8000", os.Args[0])
        }
        folder := os.Args[1]
        port := os.Args[2]
        fs := http.FileServer(http.Dir(folder))
        http.Handle("/", fs)
	log.Printf("About to listen on %s. Go to https://127.0.0.1:%s/", port, port)
	err := http.ListenAndServe(":" + port, nil)
	log.Fatal(err)
}

