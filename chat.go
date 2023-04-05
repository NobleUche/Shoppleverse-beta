package main

import (
	"log"
	"github.com/gorilla/websocket"
)


//upgrader function

var upgrader= websocket.Upgrader{
	ReadBufferSize:1024,
	WriteBufferSize:1024,
}

// Handler function

func Handler(w http.ResponseWriter, r *http.Request){
	upgrader.CheckOrigin = func(r *http.Request)bool{return true}
	conn,err:=upgrader.Upgrade(w,r,nil)
	if err!=nil{
		panic(err)
		return
	}
	for {
		messageType, p, err := conn.ReadMessage()
		if err!=nil{
			log.Println(err)
			return
		}

		err = conn.WriteMessage(messageType, p)
		if err != nil{
			log.Println(err)
			return
		}
	}
}



