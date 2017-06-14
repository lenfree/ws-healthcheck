package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/websocket"
)

type message struct {
	Action string
}

// Thanks to https://gowalker.org/golang.org/x/net/websocket#_ex_btn_Dial
func main() {
	host := os.Getenv("HOST")

	//origin := "http://" + os.Getenv("ORIGIN")
	origin := "http://localhost/"
	url := "ws://" + host + ":" + os.Getenv("PORT") + "/ws/ping"

	fmt.Printf("url: %s\n", url)
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	msg := message{Action: "ping"}
	fmt.Println("writing to ws")
	if err := websocket.JSON.Send(ws, msg); err != nil {
		log.Fatal(err)
	}

	data := message{}
	websocket.JSON.Receive(ws, &data)
	fmt.Printf("received: %#v\n", data)
}
