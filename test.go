package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	//"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	mux "github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:         ":6379",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
}

func main() {
	//genUUIDv4()
	//redisdbtest()
	//id, err := uuid.NewV4()
	//if err != nil {
	//	log.Fatalf("uuid.NewV4() failed with %s\n", err)
	//}

	redisdbtest()

	//var msgs []postmessage
	r := mux.NewRouter()
	//	msgs = append(msgs, postmessage{ID: id.String(), Message: "Test"})
	r.HandleFunc("/msgs", getMessages).Methods("GET")
	r.HandleFunc("/msgs/{id}", getMessage).Methods("GET")
	r.HandleFunc("/msgs", createMessage).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}

/* GUID-gen */
func genUUIDv4() {
	id, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("uuid.NewV4() failed with %s\n", err)
	}
	fmt.Printf("github.com/satori/go.uuid:      %s\n", id)
}

func redisdbtest() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

type postmessage struct {
	ID      string `json: "id"`
	Message string `json: "message"`
}

var msgs []postmessage
var (
	client *redis.Client
	err    error
	idTest string
)

func getMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var msgs []postmessage
	json.NewEncoder(w).Encode(msgs)
}

func getMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//var msgs []postmessage
	for _, item := range msgs {
		if item.ID == params["id"] {
			textid := client.Get(params["id"])
			file, erro := os.Create("test.txt")
			if erro != nil {
				fmt.Println("Unable to create file:", erro)
				os.Exit(1)
			}
			defer file.Close()
			file.WriteString(textid.String())
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	//textid := client.Get(params["id"])
	// file, erro := os.Create("test.txt")
	// if erro != nil {
	// 	fmt.Println("Unable to create file:", erro)
	// 	os.Exit(1)
	// }
	// defer file.Close()
	// file.WriteString(textid.String())
	json.NewEncoder(w).Encode(&postmessage{})
}

func createMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("uuid.NewV4() failed with %s\n", err)
	}

	var msg postmessage
	//var msgs []postmessage
	_ = json.NewDecoder(r.Body).Decode(&msg)
	msg.ID = id.String()
	msgs = append(msgs, msg)
	//idTest = id.String()
	client.Set(id.String(), msg.Message, time.Hour)
	json.NewEncoder(w).Encode(msg)
}
