package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
)

func Five(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprintf(w, "%d\n", 5)
}

func main() { 
    router := httprouter.New()
    router.GET("/api/five", Five)

    log.Println("Listening on 0.0.0.0:5555...")
    log.Fatal(http.ListenAndServe(":5555", router)) 
}
