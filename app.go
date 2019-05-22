package main

import(
	"fmt"
	"net/http"
	"log"
	"time"
	"encoding/json"
	"context"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"./models"
)


func home(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res,"Hey hey hey")
}

func connection() *mongo.Database{
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://srd98:Thruman98@ds211083.mlab.com:11083/teamlunch-database"))
	if err != nil{
		log.Fatal(err)
	}
	
	return client.Database("teamlunch-database")
}

func getUsers(res http.ResponseWriter, req *http.Request) {
	collection := connection().Collection("users")
	cur, err := collection.Find(context.Background(), bson.D{{}})
	users := []models.User{}
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	
	for cur.Next(context.Background()) {
			var user models.User
            err := cur.Decode(&user)
            if err != nil {
                log.Fatal(err)
            }
			users = append(users,user)
            
        }
        if err := cur.Err(); err != nil {
            log.Fatal(err)
        }
	json.NewEncoder(res).Encode(users)

}

func handleRequest(){
	
	mux := mux.NewRouter().StrictSlash(true)
	mux.HandleFunc("/",home)
	mux.HandleFunc("/users",getUsers).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081",mux))
}

func main(){
	handleRequest()
}