package services

import (
	"sync"

	"github.com/gorilla/websocket"
)

// ConnectionManager manages WebSocket connections for players.
type ConnectionManager struct {
	connections map[string]*websocket.Conn
	mu          sync.Mutex
}

// NewConnectionManager creates a new ConnectionManager.
func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		connections: make(map[string]*websocket.Conn),
	}
}

// AddConnection adds a new connection for a player.
func (cm *ConnectionManager) AddConnection(playerID string, conn *websocket.Conn) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.connections[playerID] = conn
}

// RemoveConnection removes a connection for a player.
func (cm *ConnectionManager) RemoveConnection(playerID string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	delete(cm.connections, playerID)
}

// GetConnection retrieves a connection for a player.
func (cm *ConnectionManager) GetConnection(playerID string) (*websocket.Conn, bool) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	conn, exists := cm.connections[playerID]
	return conn, exists
}
