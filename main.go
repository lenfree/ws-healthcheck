package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/websocket"
)

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
	msg := "{'action': 'ping'}"
	fmt.Println("writing to ws")
	fmt.Printf("msg: %+#v\n", msg)
	res, err := ws.Write([]byte("{'action': 'ping'}"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("response: %+#v\n", res)
}
