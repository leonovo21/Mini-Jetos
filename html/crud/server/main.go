package main

import (
	"crud/pkg/websocket"
	"fmt"
	"net/http"
)

func servWS(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("websocket endpoint reached")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		servWS(pool, w, r)
	})
}
func main() {
	fmt.Println("\u2713 Starting")
	setupRoutes()
	http.ListenAndServe(":9000", nil)
}
