package video

import (
	"sync"

	"github.com/gorilla/websocket"
)


type Participate struct{
	Host bool
	ID string
	conn *websocket.Conn
	Mutex sync.Mutex
}

type RoomMap struct{
	Mutex sync.RWMutex
	Map map[string][]Participate
}