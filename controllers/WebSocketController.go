package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSocketController struct {
	ClientId string `json:"client_id"`
	Role     string `json:"role"`
	RoomId   string `json:"room_id"`
	Action   string `json:"action"`
	Message  string `json:"message"`
}

// 连线客户端
type Client struct {
	Ws     *websocket.Conn
	UserId string
}

type Room struct {
	Id         string
	Clients    map[string]*Client
	Client_ids []string
}

var Clients = make(map[string]*Client) // 客户端列表

var Rooms = make(map[string]any) // 房间客户端列表

// Chat handles the chat page
func (s *WebSocketController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chat.html", gin.H{})
}

// WebSocketHandler handles the websocket connection
func (s *WebSocketController) Chat(c *gin.Context) {
	conn := websocket.Upgrader{
		CheckOrigin:      func(r *http.Request) bool { return true },
		HandshakeTimeout: 10 * time.Second,
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
	}
	ws, err := conn.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("upgrade:", err)
		return
	}
	defer ws.Close()

	for {
		messageType, message, err := ws.ReadMessage()

		clientId := addClient(string(message), ws) // 存入客户端信息

		if err != nil {
			fmt.Println("read:", err)
			break
		}
		switch messageType {
		case websocket.TextMessage:
			fmt.Println("text message11:", string(message))
			Clients[clientId].Ws.WriteMessage(websocket.TextMessage, []byte("message received, from client id :"+clientId))
			// ws.WriteMessage(websocket.TextMessage, message) //原始发送信息
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

// 根据信息把信息里面client存入client数组
func addClient(message string, ws *websocket.Conn) string {
	// 解析客户端信息
	clientInfo := WebSocketController{}
	err := json.Unmarshal([]byte(message), &clientInfo)
	if err != nil {
		fmt.Println("unmarshal:", err)
		return ""
	}
	// 存入客户端信息
	client := &Client{Ws: ws, UserId: clientInfo.ClientId}
	Clients[client.UserId] = client
	return client.UserId
}

/**
* mqtt server install https://docs.emqx.com/zh/emqx/latest/getting-started/getting-started.html
* mqtt client https://github.com/emqx/MQTT-Client-Examples/blob/master/mqtt-client-WebSocket/ws-mqtt.html
**/
func (s *WebSocketController) Mqtt(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "mqtt_ws.html", gin.H{})
}
