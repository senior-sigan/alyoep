package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func main() {
	log.Println("Hello from server")

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()
		log.Println("Connection!")

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				break
			}

			// Print the message to the console
			log.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			// Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				log.Println(err)
				break
			}
		}
	})

	http.Handle("/", http.FileServer(http.Dir("build/web/")))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
