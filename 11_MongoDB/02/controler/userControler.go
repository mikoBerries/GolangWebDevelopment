package controler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"m.test/11_MongoDB/02/models"
)

var (
	client   *mongo.Client
	mongoURL = "mongodb://localhost:27017"
	ctx      context.Context
	err      error
)

type UserControler struct {
	// session *mongo.Session
}

func NewUserControler() *UserControler {
	return &UserControler{}
}

func (uc UserControler) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	connectMongoDB()
	// coll := connectToDBCollection("myNewDatabase", "userData")
	s, err := client.StartSession()
	if err != nil {
		log.Panic(err.Error())
	}
	res := s.Client().Database("myNewDatabase").Collection("userData").FindOne(ctx,
		bson.D{
			{"Name", "agus"},
		})
	u := models.User{}
	res.Decode(&u)
	if err != nil {
		log.Panic(err.Error())
	}
	defer s.EndSession(ctx)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%v\n", u)
}

func (uc UserControler) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}
	//decode json to var u
	err := json.NewDecoder(r.Body).Decode(&u)
	//create bson id to data
	u.Id = primitive.NewObjectID()

	if err != nil {
		log.Panic(err.Error())
		return
	}
	connectMongoDB()
	coll := connectToDBCollection("myNewDatabase", "userData")
	s, err := client.StartSession()

	if err != nil {
		log.Panic(err.Error())
	}
	s.StartTransaction()
	//using coll from client
	res, err := coll.InsertOne(ctx, u)
	//or using session
	// res, err := s.Client().Database("myNewDatabase").Collection("userData").InsertOne(ctx, u)
	if err != nil {
		//if error
		s.EndSession(ctx)
		log.Panic(err.Error())
		return
	}
	s.CommitTransaction(ctx)
	// In MongoDB, each document stored in a collection requires a unique _id field that acts as a primary key. If an inserted document omits the _id field, the MongoDB driver automatically generates an ObjectId for the _id field.

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", res.InsertedID)
	fmt.Fprintf(w, "%v\n", u)
}

func (uc UserControler) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if id == "" {
		return
	}
	//check is "id"are object id
	if !primitive.IsValidObjectID(id) { //not a objectid
		log.Panic("invalid")
		return
	}
	s, err := client.StartSession()
	if err != nil {
		log.Panic(err.Error())
	}
	oid, _ := primitive.ObjectIDFromHex(id)

	//filter ==> while clause
	filter := bson.D{{"_id", oid}}
	_, err = s.Client().Database("myNewDatabase").Collection("userData").DeleteOne(ctx,
		filter)
	if err != nil {
		panic(err)
	}
	w.Header().Set("content-type", "text/html")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "some magic code to delete user")
}

func connectMongoDB() {

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Panic(err.Error())
	}
	// ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	// defer cancel()

	ctx, _ = context.WithTimeout(context.Background(), 20*time.Second)
	client.Connect(ctx) //client will have time out 20 sec

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Panic(err.Error())
	}
	// coll := client.Database("myNewDatabase").Collection("userData")
	// res := coll.FindOne(ctx,
	// 	bson.D{
	// 		{"Name", "agus"},
	// 	})
	// u := models.User{}
	// res.Decode(&u)
	// fmt.Println(u)
}

func connectToDBCollection(d string, c string) *mongo.Collection {

	coll := client.Database(d).Collection(c)

	return coll
}
