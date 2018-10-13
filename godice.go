package main

import (
  "fmt"
  "html/template"
  "log"
  "math/rand"
  "net/http"
)

type Dice struct {
  Roll int
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func diceHandler(w http.ResponseWriter, r *http.Request) {
  d := Dice{Roll: rand.Intn(6)}
  t, _ := template.ParseFiles("dice.html")
  t.Execute(w, d)
}

func main() {
  http.HandleFunc("/dice", diceHandler)
  http.HandleFunc("/", rootHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
