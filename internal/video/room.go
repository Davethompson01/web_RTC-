package video

import (
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Participant struct {
	Host  bool
	ID    string
	conn  *websocket.Conn
	Mutex sync.Mutex
}

type RoomMap struct {
	Mutex sync.RWMutex
	Map   map[string][]Participant
}

func (r *RoomMap) Init() {
	r.Map = make(map[string][]Participant)
}

func (r *RoomMap) Get(roomID string) []Participant {
	r.Mutex.RLock()

	defer r.Mutex.RUnlock()
	return r.Map[roomID]
}

func (r *RoomMap) CreateRoom() string {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	rand.New(rand.NewSource(time.Now().UnixNano()))
	var letters = []rune(
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789",
	)

	b := make([]rune, 8)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	roomID := string(b)

	r.Map[roomID] = []Participant{}
	return roomID
}

func (r *RoomMap) InsertIntoRoom(
	Host bool,
	roomID string,
	conn *websocket.Conn,
) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	clientID := uuid.New().String()
	incomingParticipant := Participant{
		Host, clientID, conn, sync.Mutex{},
	}

	r.Map[roomID] = append(r.Map[roomID], incomingParticipant)
}



func (r *RoomMap) delete(roomID string){
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	delete(r.Map, roomID)
}