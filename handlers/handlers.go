package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"../models"
	Conf "../redisconf"
	Mux "github.com/gorilla/mux"
	Uuid "github.com/satori/go.uuid"
)

//Msgs init
var Msgs []models.Postmessage
var (
	//Err error messages
	Err error
)

//Handlers func interface
type Handlers interface {
	GetMessage(W http.ResponseWriter, R *http.Request)
	GetMessages(W http.ResponseWriter, R *http.Request)
	CreateMessage(W http.ResponseWriter, R *http.Request)
}

//GetMessages read all messages from buffer
func GetMessages(W http.ResponseWriter, R *http.Request) {
	W.Header().Set("Content-Type", "application/json")
	json.NewEncoder(W).Encode(Msgs)
}

//GetMessage read message from buffer and Redis by uid
func GetMessage(W http.ResponseWriter, R *http.Request) {
	W.Header().Set("Content-Type", "application/json")
	Params := Mux.Vars(R)
	for _, Item := range Msgs {
		if Item.ID == Params["id"] {
			Textid := Conf.Client.Get(Params["id"])
			File, Erro := os.Create("test.txt")
			if Erro != nil {
				fmt.Println("Unable to create file:", Erro)
				os.Exit(1)
			}
			defer File.Close()
			File.WriteString(Textid.String())
			json.NewEncoder(W).Encode(Item)
			return
		}
	}
	json.NewEncoder(W).Encode(&models.Postmessage{})
}

//CreateMessage create message in buffer and Redis
func CreateMessage(W http.ResponseWriter, R *http.Request) {
	W.Header().Set("Content-Type", "application/json")
	id, Err := Uuid.NewV4()
	if Err != nil {
		log.Fatalf("uuid.NewV4() failed with %s\n", Err)
	}
	var Msg models.Postmessage
	_ = json.NewDecoder(R.Body).Decode(&Msg)
	Msg.ID = id.String()
	Msgs = append(Msgs, Msg)
	Err1 := Conf.Client.Set(Msg.ID, Msg.Message, 0).Err()
	if Err1 != nil {
		panic(Err1)
	}
	json.NewEncoder(W).Encode(Msg)
}
