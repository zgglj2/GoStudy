package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:9090", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
	fmt.Printf("🔗 Connecting to WebSocket server: %s\n", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("❌ Dial error:", err)
	}
	defer c.Close()
	fmt.Println("✅ WebSocket connected!")

	done := make(chan struct{})

	// Read messages goroutine
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				fmt.Println("❌ Read error:", err)
				return
			}
			fmt.Printf("📨 Received: %s\n", message)
		}
	}()

	// User input handling goroutine
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("\n💡 Usage instructions:")
		fmt.Println("   - Type any text to send a message")
		fmt.Println("   - Type 'ping' to send a ping test")
		fmt.Println("   - Type 'quit' or press Ctrl+C to exit")
		fmt.Println("   - Type 'auto' to start auto sending test messages")
		fmt.Print("\n> ")

		for scanner.Scan() {
			text := strings.TrimSpace(scanner.Text())
			
			switch text {
			case "quit", "exit":
				fmt.Println("👋 User initiated quit")
				interrupt <- os.Interrupt
				return
			case "ping":
				err := c.WriteMessage(websocket.PingMessage, []byte("ping"))
				if err != nil {
					fmt.Println("❌ Ping send error:", err)
					return
				}
				fmt.Println("🏓 Ping sent")
			case "auto":
				fmt.Println("🤖 Starting auto send test messages...")
				go autoSendMessages(c)
			case "":
				// Empty input, ignore
			default:
				err := c.WriteMessage(websocket.TextMessage, []byte(text))
				if err != nil {
					fmt.Println("❌ Send message error:", err)
					return
				}
				fmt.Printf("📤 Sent: %s\n", text)
			}
			fmt.Print("> ")
		}
	}()

	// Main loop
	for {
		select {
		case <-done:
			return
		case <-interrupt:
			fmt.Println("\n🛑 Interrupt received, closing connection...")
			
			// Send close message
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				fmt.Println("❌ Close message send error:", err)
				return
			}
			
			// Wait for connection to close
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			fmt.Println("✅ Connection closed")
			return
		}
	}
}

func autoSendMessages(c *websocket.Conn) {
	messages := []string{
		"Hello WebSocket!",
		"This is a test message",
		"Testing special chars: !@#$%^&*()",
		"JSON test: {\"name\":\"test\",\"value\":123}",
		"Long message: " + strings.Repeat("A", 100),
	}

	for i, msg := range messages {
		time.Sleep(2 * time.Second)
		fullMsg := fmt.Sprintf("[Auto %d] %s", i+1, msg)
		err := c.WriteMessage(websocket.TextMessage, []byte(fullMsg))
		if err != nil {
			fmt.Println("❌ Auto send error:", err)
			return
		}
		fmt.Printf("🤖 Auto sent: %s\n> ", fullMsg)
	}
	fmt.Println("🤖 Auto send completed")
	fmt.Print("> ")
}