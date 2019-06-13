package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

  api "github.com/shreddedbacon/thingiverse-go"
)

func (ta *ThingiApp) DisplayNewest(w http.ResponseWriter, r *http.Request) {
  result, err := ta.thingiAPI.GetNewest()
  if err != nil {
    log.Println(err)
  }
  thingiThings := api.Things{}
  json.Unmarshal([]byte(result), &thingiThings)
  tmpl, err := template.New("").ParseFiles("templates/newest.html", "templates/base.html")
  // check your err
  err = tmpl.ExecuteTemplate(w, "base", thingiThings)
  if err != nil{
    log.Println(err)
  }
}

func (ta *ThingiApp) DisplayPopular(w http.ResponseWriter, r *http.Request) {
  result, err := ta.thingiAPI.GetPopular()
  if err != nil {
    log.Println(err)
  }
  thingiThings := api.Things{}
  json.Unmarshal([]byte(result), &thingiThings)
  tmpl, err := template.New("").ParseFiles("templates/popular.html", "templates/base.html")
  // check your err
  err = tmpl.ExecuteTemplate(w, "base", thingiThings)
  if err != nil{
    log.Println(err)
  }
}

func (ta *ThingiApp) DisplayFeatured(w http.ResponseWriter, r *http.Request) {
  result, err := ta.thingiAPI.GetFeatured(false)
  if err != nil {
    log.Println(err)
  }
  thingiThings := api.Things{}
  json.Unmarshal([]byte(result), &thingiThings)
  tmpl, err := template.New("").ParseFiles("templates/featured.html", "templates/base.html")
  // check your err
  err = tmpl.ExecuteTemplate(w, "base", thingiThings)
  if err != nil{
    log.Println(err)
  }
}
