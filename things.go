package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

  api "github.com/shreddedbacon/thingiverse-go"
  "github.com/gorilla/mux"
)

func (ta *ThingiApp) DisplayThing(w http.ResponseWriter, r *http.Request) {
  urlvars := mux.Vars(r)
  thingId := urlvars["thingId"]
  result, err := ta.thingiAPI.GetThing(thingId)
  if err != nil {
    log.Println(err)
  }
  thingiThings := api.Thing{}
  json.Unmarshal([]byte(result), &thingiThings)
  tmpl, err := template.New("").ParseFiles("templates/thing.html", "templates/base.html")
  // check your err
  err = tmpl.ExecuteTemplate(w, "base", thingiThings)
  if err != nil{
    log.Println(err)
  }
}
