package ws

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// ConnectionManager manages a persistent WebSocket connection
type ConnectionManager struct {
	conn         *websocket.Conn
	latestMsg    []byte
	msgType      int
	mu           sync.RWMutex
	done         chan struct{}
	reconnecting bool
	url          string
}

var (
	manager *ConnectionManager
	once    sync.Once
)

// GetManager returns the singleton connection manager
func GetManager() *ConnectionManager {
	once.Do(func() {
		manager = &ConnectionManager{
			done: make(chan struct{}),
			url:  "ws://localhost:42069/ws",
		}
	})
	return manager
}

// Connect establishes the WebSocket connection and starts the reader
func (cm *ConnectionManager) Connect() error {
	if cm.conn != nil {
		return nil // Already connected
	}

	conn, _, err := websocket.DefaultDialer.Dial(cm.url, nil)
	if err != nil {
		return err
	}

	cm.conn = conn
	cm.done = make(chan struct{})

	// Start background reader
	go cm.readLoop()

	log.Println("WebSocket connected")
	return nil
}

// readLoop continuously reads messages in the background
func (cm *ConnectionManager) readLoop() {
	defer close(cm.done)

	for {
		msgType, msg, err := cm.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket error: %v", err)
			}
			cm.handleDisconnect()
			return
		}

		// Store latest message thread-safely
		cm.mu.Lock()
		cm.msgType = msgType
		cm.latestMsg = msg
		cm.mu.Unlock()

		log.Printf("Received: %s", string(msg))
	}
}

// handleDisconnect manages reconnection logic
func (cm *ConnectionManager) handleDisconnect() {
	cm.mu.Lock()
	cm.conn = nil
	cm.reconnecting = true
	cm.mu.Unlock()

	log.Println("WebSocket disconnected, attempting to reconnect...")

	// Try to reconnect with exponential backoff
	backoff := time.Second
	maxBackoff := 30 * time.Second

	for {
		time.Sleep(backoff)

		err := cm.Connect()
		if err == nil {
			cm.mu.Lock()
			cm.reconnecting = false
			cm.mu.Unlock()
			log.Println("WebSocket reconnected")
			return
		}

		log.Printf("Reconnection failed: %v", err)
		if backoff < maxBackoff {
			backoff *= 2
		}
	}
}

// GetLatestMessage returns the most recent message received (thread-safe)
func (cm *ConnectionManager) GetLatestMessage() (int, []byte) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	return cm.msgType, cm.latestMsg
}

// SendMessage sends a message to the server
func (cm *ConnectionManager) SendMessage(msgType int, data []byte) error {
	cm.mu.RLock()
	conn := cm.conn
	cm.mu.RUnlock()

	if conn == nil {
		return nil // Silently drop if not connected
	}

	return conn.WriteMessage(msgType, data)
}

// IsConnected returns true if the connection is active
func (cm *ConnectionManager) IsConnected() bool {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	return cm.conn != nil && !cm.reconnecting
}

// Close closes the WebSocket connection
func (cm *ConnectionManager) Close() {
	cm.mu.Lock()
	conn := cm.conn
	cm.conn = nil
	cm.mu.Unlock()

	if conn != nil {
		conn.Close()
		<-cm.done // Wait for readLoop to finish
	}
}

// Legacy function for backward compatibility
func ConnectToWS() {
	mgr := GetManager()
	if err := mgr.Connect(); err != nil {
		log.Fatal("Failed to connect to WebSocket:", err)
	}
}

// GetMessage returns the latest message (for use in game loop)
func GetMessage() (int, []byte) {
	return GetManager().GetLatestMessage()
}

func ParseLatestMessage() (string, string) {
	_, msg := GetMessage()
	if msg == nil {
		return "", ""
	}
	var msgJSON Message

	err := json.Unmarshal(msg, &msgJSON)
	if err != nil {
		log.Fatal("error pasrsing json: ", err)
	}
	return msgJSON.Type, msgJSON.Payload
}
