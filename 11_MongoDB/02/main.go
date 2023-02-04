package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"m.test/11_MongoDB/02/controler"
)

func main() {

	uc := controler.NewUserControler()
	r := httprouter.New()
	// r.GET("/", index)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/createUser/", uc.CreateUser)
	r.DELETE("/deleteUser", uc.DeleteUser)
	log.Panic(http.ListenAndServe(":8080", r))

}

// func getSession() *mongo.Session {
// 	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
// 	defer cancel()
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://foo:bar@localhost:27017"))
// 	if err != nil {
// 		fmt.Print(err.Error())
// 	}
// 	s, err := client.StartSession()
// 	if err != nil {
// 		fmt.Print(err.Error())
// 	}
// 	err = client.Ping(ctx, readpref.Primary())
// 	if err != nil {
// 		fmt.Print(err.Error())
// 	}
// 	return &s
// }
