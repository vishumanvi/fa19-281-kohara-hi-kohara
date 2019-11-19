/*
	Follow Post Microservice
*/

package main

import (
	"encoding/json"
	"log"
	"os"

	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// MongoDB Config
var mongodb_server = "mongodb://localhost:32768/?authSource=admin"
var mongodb_database = "like"
var mongodb_collection = "like"

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	if len(os.Getenv("MONGO")) == 0 {
		mongodb_server = os.Getenv("MONGO")
	}
	return n
}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/{id}", addNewLikeHandler(formatter)).Methods("POST")
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Like Write alive!"})
	}
}

// API to add a new like
func addNewLikeHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		params := mux.Vars(req)
		postId := params["id"]
		var requestBody Like
		_ = json.NewDecoder(req.Body).Decode(&requestBody)
		create := bson.M{
			"Username": requestBody.Username,
			"PostId":   postId}
		err = c.Insert(create)
		if err != nil {
			log.Fatal(err)
		}
		formatter.JSON(w, http.StatusOK, requestBody.Username)
	}
}
