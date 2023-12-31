package main

import (
	"net/http"

	"github.com/CrimsonCoder42/mongo-golang/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

// Older version uses gopkg but mongo has it's own package https://www.mongodb.com/docs/drivers/go/current/

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())

	r.GET("/user/:id", uc.GetUser)

	r.POST("/user", uc.CreateUser)

	r.DELETE("/user/:id", uc.RemoveUser)

	http.ListenAndServe("localhost:8080", r)

}

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost:27017")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}
