package routers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// WebSocketHandler handles the websocket connection
func WebSocketHandler(c *gin.Context) {
	wsUpgrader := websocket.Upgrader{
		CheckOrigin:      func(r *http.Request) bool { return true },
		HandshakeTimeout: 10 * time.Second,
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
	}
	ws, err := wsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("upgrade:", err)
		return
	}
	defer ws.Close()

	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
			break
		}

		switch messageType {
		case websocket.TextMessage:
			fmt.Println("text message11:", string(message))
			ws.WriteMessage(websocket.TextMessage, message)
		case websocket.BinaryMessage:
			fmt.Println("this is file type")
		case websocket.CloseMessage:
			fmt.Println("close websocket")
		case websocket.PingMessage:
			ws.WriteMessage(websocket.TextMessage, []byte("ping"))
			fmt.Println("this is ping message")
		case websocket.PongMessage:
			ws.WriteMessage(websocket.TextMessage, []byte("pong"))
			fmt.Println("this is pong message")
		default:
			fmt.Println("unknown message type:", messageType)
		}
	}

}
