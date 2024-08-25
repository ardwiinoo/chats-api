package ws

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	hub *Hub
}

type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewHandler(h *Hub) *Handler {
	return &Handler{
		hub: h,
	}
}

func (h *Handler) CreateRoom(gctx *gin.Context) {
	var payload CreateRoomReq

	if err := gctx.ShouldBindJSON(&payload); err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.hub.Rooms[payload.ID] = &Room{
		ID: payload.ID,
		Name: payload.Name,
		Clients: make(map[string]*Client),
	}

	gctx.JSON(http.StatusOK, payload)
}

// gorilla websocket
var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // can be front-end specific domain
	},
}

func (h *Handler) JoinRoom(gctx *gin.Context) {
	conn, err := upgrader.Upgrade(gctx.Writer, gctx.Request, nil)
	if err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roomID := gctx.Param("roomId")
	clientID := gctx.Query("userId")
	username := gctx.Query("username")

	cl := &Client{
		Conn: conn,
		Message: make(chan *Message, 10),
		ID: clientID,
		RoomID: roomID,
		Username: username,
	}

	msg := &Message{
		Content: "A new user has join the room",
		RoomID: roomID,
		Username: username,
	}

	// reg new client & broadcast msg
	h.hub.Register <- cl
	h.hub.Broadcast <- msg

	go cl.writeMessage()
	cl.readMessage(h.hub)
}

