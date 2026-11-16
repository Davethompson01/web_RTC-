package video

import "github.com/gorilla/websocket"

var AllRooms RoomMap

type response struct {
	RoomID string `json:"roomID"`
}

type BroadcastMessage struct{
	Message map[string]interface{}
	RoomID string
	Client *websocket.Conn
}

