package client

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"log"
	"net"
)

func HandleConnection(connection net.Conn) {

	hostAddress := connection.RemoteAddr().String()

	buffer := make([]byte, 2048)
	size, err := connection.Read(buffer)
	if err != nil {
		log.Println("Failed to read data from host " + hostAddress + ": " + err.Error())
		return
	}
	payload := string(buffer[:size])
	log.Println("Recieved Data from host " + hostAddress)
	if !gjson.Valid(payload) {
		log.Println("Recieved data from host " + hostAddress + "was not json")
		_ = connection.Close()
		return
	}

	data := Response{
		ResponseType: "authentication",
		ResponseData: map[string]string{
			"success":           "true",
			"authenticationkey": "testKey",
			"nickname":          "nicknameTest",
		},
	}
	response := make([]byte, 2048)
	response, err = json.Marshal(&data)
	if err != nil {
		log.Println("error marshalling response on host " + hostAddress + ": " + err.Error())
		_ = connection.Close()
		return
	}
	_, err = connection.Write(response)
	if err != nil {
		log.Println("error writing data back to host " + hostAddress + ": " + err.Error())
		_ = connection.Close()
		return
	}
	log.Println("Responded with data " + string(response))

}

type Response struct {
	ResponseType string            `json:"type"`
	ResponseData map[string]string `json:"data"`
}
