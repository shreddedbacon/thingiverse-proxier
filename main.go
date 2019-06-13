package main

import (
	"log"
	"net/http"
	"os"
  "fmt"

  api "github.com/shreddedbacon/thingiverse-go"
  "github.com/gorilla/mux"
)

type thingiApp interface {
  DisplayNewest(w http.ResponseWriter, r *http.Request)
  DisplayPopular(w http.ResponseWriter, r *http.Request)
  DisplayFeatured(w http.ResponseWriter, r *http.Request)
  DisplayThing(w http.ResponseWriter, r *http.Request)
  HomeHandler(w http.ResponseWriter, r *http.Request)
}

type ThingiApp struct {
  thingiAPI api.TthingiAPI
}

func NewThingiApp(thingiAPI api.TthingiAPI) thingiApp {
	return &ThingiApp{
		thingiAPI:     thingiAPI,
	}
}

func main() {
	if len(os.Getenv("CLIENTID")) == 0 {
		log.Fatalln("CLIENTID not set")
	}
	if len(os.Getenv("CLIENTSECRET")) == 0 {
		log.Fatalln("CLIENTSECRET not set")
	}
	if len(os.Getenv("APPTOKEN")) == 0 {
		log.Fatalln("APPTOKEN not set")
	}

	clientId := os.Getenv("CLIENTID")
	clientSecret := os.Getenv("CLIENTSECRET")
	appToken := os.Getenv("APPTOKEN")

	thingiAPI := api.NewAPI(clientId, clientSecret, appToken)
  thingiApp := NewThingiApp(thingiAPI)

  r := mux.NewRouter()


  r.Path("/").Methods(http.MethodGet).HandlerFunc(thingiApp.HomeHandler)
	r.Path("/newest").Methods(http.MethodGet).HandlerFunc(thingiApp.DisplayNewest)
  r.Path("/popular").Methods(http.MethodGet).HandlerFunc(thingiApp.DisplayPopular)
  r.Path("/featured").Methods(http.MethodGet).HandlerFunc(thingiApp.DisplayFeatured)
  r.Path("/thing/{thingId}").Methods(http.MethodGet).HandlerFunc(thingiApp.DisplayThing)

	http.ListenAndServe(":8900", r)
}

func (ta *ThingiApp) HomeHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, `<h1>Thingiverse Proxier</h1>
    <a href="/newest">newest</a><br>
    <a href="/popular">popular</a><br>
    <a href="/featured">featured</a>`)
}
