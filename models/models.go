package models

//Postmessage model for Handlers
type Postmessage struct {
	ID      string `json: "id"`
	Message string `json: "message"`
}
