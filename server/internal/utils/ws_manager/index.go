package ws_manager

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
}

func (c *Client) Monitor() {
	go func() {
		for {
			// msgType, data, err := c.Conn.ReadMessage()
			// if err != nil {
			// 	break
			// }
		}
	}()
}

type ClientManager struct {
	clients map[string]*Client
	*sync.RWMutex
}

var ClientMgr = ClientManager{
	clients: make(map[string]*Client),
	RWMutex: &sync.RWMutex{},
}

func (cm *ClientManager) Run() {
}

func (cm *ClientManager) AddClient(clientId string, conn *websocket.Conn) {
	cm.Lock()
	defer cm.Unlock()

	cm.clients[clientId] = &Client{
		Conn: conn,
	}
}

func (cm *ClientManager) RemoveClient(clientId string) {
	cm.Lock()
	defer cm.Unlock()

	delete(cm.clients, clientId)
}

func (cm *ClientManager) GetClients() map[string]*Client {
	cm.RLock()
	defer cm.RUnlock()

	return cm.clients
}
