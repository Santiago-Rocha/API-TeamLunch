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
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func getTurn(res http.ResponseWriter, req *http.Request) {
	//key, err := primitive.ObjectIDFromHex("5c3ec39b41b83d1d740491a6")
	keys := mapID([]string{"5c3ec39b41b83d1d740491a6","5c42078cfb4b960017716918"});
	pipeline := []bson.M{ 
		bson.M{
			"$match": bson.M{
				"_id": bson.M{"$in":keys},
			},
		},
		bson.M{
			"$group": bson.M{
				"_id": bson.M{ 
					"benefits" : bson.M{ "$multiply" : bson.A{"$benefits",-1} }, 
					"date": "$last_lunch"},
				"docs": bson.M{ "$push" : "$$ROOT"},
			},
		},
		bson.M{"$sort": bson.M{"_id": 1 }},
		bson.M{"$limit": 2},
	}
	
	collection := connection().Collection("users")
	opts := options.Aggregate()
	cur, err := collection.Aggregate(context.Background(), pipeline, opts); 
	resp := []bson.M{}
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var set bson.M
        err := cur.Decode(&set)
        if err != nil {
            log.Fatal(err)
        }
		resp = append(resp,set)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(res).Encode(resp)	
}

func mapID(IDs []string) bson.A {
	var result bson.A
	for _, data := range IDs {
		d, _ := primitive.ObjectIDFromHex(data)
		result = append(result,d)
	}
	fmt.Println(result)
	return result
} 

func handleRequest(){
	
	mux := mux.NewRouter().StrictSlash(true)
	mux.HandleFunc("/",home)
	mux.HandleFunc("/users",getUsers).Methods("GET")
	mux.HandleFunc("/turn",getTurn).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081",mux))
}

func main(){
	handleRequest()
}