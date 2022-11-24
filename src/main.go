package main

import (
	"InternetRelayChat-Backend/src/client"
	"InternetRelayChat-Backend/src/database"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "8888"
	SERVER_TYPE = "tcp4"
)

func init() {
	createFolders()
}

func main() {

	database.Init()
	defer database.Close()

	log.Println("Starting Server...")
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		log.Fatalln("Error attempting to listen: " + err.Error())
	}
	defer server.Close()
	log.Println("Started Server Successfully!")
	log.Println("Listening at " + SERVER_HOST + ":" + SERVER_PORT)

	go func() {
		fmt.Println("Awaiting Connections...")
		for {
			connection, err := server.Accept()
			if err != nil {
				fmt.Println("Error accepting: ", err.Error())
				os.Exit(1)
			}
			fmt.Println("Client connected at " + connection.RemoteAddr().String())
			go client.HandleConnection(connection)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("\u001B[38:5:1mPress Ctrl+C to exit")
	<-stop

}

func createFolders() {
	{
		err := os.Mkdir("data", 0750)
		if err != nil && !errors.Is(err, os.ErrExist) {
			log.Println(err)
		}
	}
}
