package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"golang.org/x/net/websocket"
)

// Thanks to https://gowalker.org/golang.org/x/net/websocket#_ex_btn_Dial
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("URL")
	origin := "ws://" + os.Getenv("ORIGIN") + ":" + os.Getenv("PORT") + "/ws/ping"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := ws.Write([]byte("hello, world!\n")); err != nil {
		log.Fatal(err)
	}
	var msg = make([]byte, 512)
	var n int
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Received: %s.\n", msg[:n])
}
