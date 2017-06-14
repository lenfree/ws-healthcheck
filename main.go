package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/websocket"
)

type message struct {
	action string
}

type resTime struct {
	Time string
}

type response struct {
	Action string
	Req    message
	Res    resTime
}

// Thanks to https://gowalker.org/golang.org/x/net/websocket#_ex_btn_Dial
func main() {
	host := os.Getenv("HOST")

	//origin := "http://" + os.Getenv("ORIGIN")
	origin := "http://localhost/"
	url := "ws://" + host + ":" + os.Getenv("PORT") + "/ws"

	fmt.Printf("url: %s\n", url)
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("send request")
	msg := message{action: "PING"}
	fmt.Println("writing to ws")
	fmt.Printf("msg: %+#v\n", msg)
	if err := websocket.JSON.Send(ws, msg); err != nil {
		log.Fatal(err)
	}

	fmt.Println("read response")
	var data interface{}
	websocket.JSON.Receive(ws, &data)
	fmt.Printf("received: %#v\n", data)
}
