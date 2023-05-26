package router

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/FulgurCode/school-erp-api/helpers"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Message struct {
	TargetID string      `json:"targetId"`
	Name     string      `json:"name"`
	Data     interface{} `json:"data"`
}

var connections = map[string]*websocket.Conn{}
var connectionsM sync.Mutex

func HandleWebSocket(c *gin.Context) {
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// Allow all connections
			return true
		},
	}
	// Getting socket connection
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		// Handle error
		return
	}
	// Generating unique id and storing in connections
	var id = helpers.GenerateUniqueID()
	connectionsM.Lock()
	connections[id] = conn
	connectionsM.Unlock()
	conn.WriteJSON(map[string]interface{}{"name": "id", "data": id})
	// Handling socket messages
	for {
		// Reading and decoding message from client
		_, msg, err := conn.ReadMessage()
		if err != nil {
			// Handle error or connection close
			break
		}
		var data Message
		json.Unmarshal(msg, &data)
		// Getting connection using target id
		targetID := data.TargetID
		connectionsM.Lock()
		targetConn, exists := connections[targetID]
		connectionsM.Unlock()
		if exists {
			targetConn.WriteJSON(map[string]interface{}{"name": data.Name, "data": data.Data})
		}
	}
	// Closing connection and removing from connections
	conn.Close()
	connectionsM.Lock()
	delete(connections, id)
	connectionsM.Unlock()
}
