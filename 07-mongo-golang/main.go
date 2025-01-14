package main

import (
	"fmt"
	"mongo-golang/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe(":8000", r)
}

func getSession() *mgo.Session {
	fmt.Println("starting mongodb connection")
	s, err := mgo.Dial("mongodb://myuser:mypassword@127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to mongodb")
	return s
}
