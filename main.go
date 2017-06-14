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

	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	msg := []byte("{\"action\": \"PING\"}")
	_, err = ws.Write(msg)
	if err != nil {
		log.Fatal(err)
	}

	var res = make([]byte, 4096)
	var n int
	if n, err = ws.Read(res); err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(res[:n]))
}
