package utils

import (
	"encoding/json"
	"log"
)

type Message struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func MarshalJsonMessage(name string, msg string) ([]byte, error) {
	newServerMessage := Message{
		Name:    name,
		Message: msg,
	}
	jsonMessage, err := json.Marshal(newServerMessage)

	if err != nil {
		log.Println("marshal json error")
		return nil, err
	}

	return jsonMessage, nil
}
