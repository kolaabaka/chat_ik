package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var checkOrigin = func(r *http.Request) bool {
	secret := r.Header.Get("secret")
	if len(secret) == 0 {
		return false
	}

	return true
}

var upgrader = websocket.Upgrader{
	CheckOrigin: checkOrigin,
}

func webSocketHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	defer ws.Close()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("read error:", err)
			break
		}
		fmt.Printf("Received: %s\n", msg)

		if err := ws.WriteMessage(websocket.TextMessage, msg); err != nil {
			fmt.Println("write error:", err)
			break
		}
	}
}
