package main

import (
	"code.google.com/p/go.net/websocket"
	//"io"
	"log"
	"net/http"
)

func echoHandler(ws *websocket.Conn) {
	type T struct {
		X     int64
		Y     int64
		Color string
		Str   string
	}

	var data T
	websocket.JSON.Receive(ws, &data)

	log.Printf("%#v\n", data)

	websocket.JSON.Send(ws, data)
}

func main() {
	http.Handle("/echo", websocket.Handler(echoHandler))
	http.Handle("/", http.FileServer(http.Dir("./")))
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
