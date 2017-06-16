package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/websocket"
)

// Thanks to https://gowalker.org/golang.org/x/net/websocket#_ex_btn_Dial
func main() {
	host := os.Getenv("HOST")
	if len(host) <= 0 {
		log.Fatalf("%s\n", "Please specify host endpoint")
	}

	var origin string
	origin = "http://" + os.Getenv("ORIGIN")
	if len(origin) <= 0 {
		origin = "http://localhost/"
	}

	url := "ws://" + host + ":" + os.Getenv("PORT") + "/ws"

	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatalf("Connect to websocket error %s\n", err)
		os.Exit(2)
	}

	msg := []byte("{\"action\": \"PING\"}")
	_, err = ws.Write(msg)
	if err != nil {
		log.Fatalf("Writing to websocket error %s\n", err)
		os.Exit(2)
	}

	var res = make([]byte, 4096)
	var n int
	if n, err = ws.Read(res); err != nil {
		log.Fatalf("Reading from websocket error %s\n", err)
		os.Exit(2)
	}

	data := "\"action\":\"PONG\""
	if strings.Contains(string(res[:n]), data) {
		fmt.Println(string(res[:n]))
	} else {
		log.Fatalf("Response did not contain %s", data)
	}
}
