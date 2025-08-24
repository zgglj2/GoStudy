package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:9090", "http service address")

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins, should be more strict in production
	},
}

// Connection statistics
var stats struct {
	connections int
	messages    int
}

func echo(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("❌ Upgrade failed: ", err)
		return
	}
	defer conn.Close()

	stats.connections++
	clientAddr := conn.RemoteAddr().String()
	fmt.Printf("✅ New connection: %s (Total connections: %d)\n", clientAddr, stats.connections)

	// Set ping/pong handlers
	conn.SetPingHandler(func(appData string) error {
		fmt.Printf("🏓 Received ping from %s: %s\n", clientAddr, appData)
		return conn.WriteMessage(websocket.PongMessage, []byte(appData))
	})

	conn.SetPongHandler(func(appData string) error {
		fmt.Printf("🏓 Received pong from %s: %s\n", clientAddr, appData)
		return nil
	})

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("❌ Connection unexpectedly closed %s: %v\n", clientAddr, err)
			} else {
				fmt.Printf("🔌 Connection closed normally: %s\n", clientAddr)
			}
			break
		}

		stats.messages++

		// Handle based on message type
		switch mt {
		case websocket.TextMessage:
			fmt.Printf("📨 Received text message from %s: %s\n", clientAddr, message)

			// Special command handling
			msgStr := string(message)
			switch msgStr {
			case "time":
				response := fmt.Sprintf("Server time: %s", time.Now().Format("2006-01-02 15:04:05"))
				err = conn.WriteMessage(websocket.TextMessage, []byte(response))
			case "stats":
				response := fmt.Sprintf("Server stats - Connections: %d, Messages: %d", stats.connections, stats.messages)
				err = conn.WriteMessage(websocket.TextMessage, []byte(response))
			case "ping":
				err = conn.WriteMessage(websocket.PingMessage, []byte("server-ping"))
			default:
				// Echo message
				response := fmt.Sprintf("Echo: %s", message)
				err = conn.WriteMessage(mt, []byte(response))
			}

		case websocket.BinaryMessage:
			fmt.Printf("📦 Received binary message from %s: %d bytes\n", clientAddr, len(message))
			err = conn.WriteMessage(mt, message)

		case websocket.PingMessage:
			fmt.Printf("🏓 Received ping message from %s\n", clientAddr)
			err = conn.WriteMessage(websocket.PongMessage, message)

		case websocket.PongMessage:
			fmt.Printf("🏓 Received pong message from %s\n", clientAddr)
		}

		if err != nil {
			fmt.Printf("❌ Send message error to %s: %v\n", clientAddr, err)
			break
		}
	}

	stats.connections--
	fmt.Printf("👋 Connection disconnected: %s (Remaining connections: %d)\n", clientAddr, stats.connections)
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

func status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"connections": %d, "messages": %d, "uptime": "%s"}`,
		stats.connections, stats.messages, time.Since(startTime))
}

var startTime = time.Now()

func main() {
	flag.Parse()
	log.SetFlags(0)

	fmt.Printf("🚀 Starting WebSocket debug server\n")
	fmt.Printf("📍 Address: %s\n", *addr)
	fmt.Printf("🌐 Web interface: http://%s\n", *addr)
	fmt.Printf("📊 Status endpoint: http://%s/status\n", *addr)
	fmt.Printf("🔗 WebSocket endpoint: ws://%s/echo\n", *addr)

	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)
	http.HandleFunc("/status", status)

	fmt.Println("✅ Server started successfully, waiting for connections...")
	log.Fatal(http.ListenAndServe(*addr, nil))
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>WebSocket Debug Tool</title>
<style>
body { font-family: Arial, sans-serif; margin: 20px; }
.container { max-width: 800px; margin: 0 auto; }
.status { background: #f0f0f0; padding: 10px; border-radius: 5px; margin: 10px 0; }
.messages { height: 300px; overflow-y: scroll; border: 1px solid #ccc; padding: 10px; background: #f9f9f9; }
.input-group { margin: 10px 0; }
input[type="text"] { width: 70%; padding: 5px; }
button { padding: 5px 10px; margin: 0 5px; }
.connected { color: green; }
.disconnected { color: red; }
</style>
</head>
<body>
<div class="container">
<h1>🔧 WebSocket Debug Tool</h1>
<div class="status">
    <p>Connection status: <span id="status" class="disconnected">Disconnected</span></p>
    <p>WebSocket address: <code>{{.}}</code></p>
</div>

<div class="input-group">
    <button id="connect">Connect</button>
    <button id="disconnect">Disconnect</button>
    <button id="clear">Clear Log</button>
</div>

<div class="input-group">
    <input id="messageInput" type="text" placeholder="Enter message..." value="Hello WebSocket!">
    <button id="send">Send</button>
    <button id="sendTime">Get Time</button>
    <button id="sendStats">Get Stats</button>
    <button id="sendPing">Send Ping</button>
</div>

<h3>📋 Message Log</h3>
<div id="messages" class="messages"></div>
</div>

<script>
var ws;
var output = document.getElementById("messages");
var status = document.getElementById("status");
var messageInput = document.getElementById("messageInput");

function log(message, type = 'info') {
    var d = document.createElement("div");
    var time = new Date().toLocaleTimeString();
    d.innerHTML = '<span style="color: #666;">[' + time + ']</span> ' + message;
    if (type === 'error') d.style.color = 'red';
    if (type === 'success') d.style.color = 'green';
    if (type === 'warning') d.style.color = 'orange';
    output.appendChild(d);
    output.scrollTop = output.scrollHeight;
}

document.getElementById("connect").onclick = function() {
    if (ws) return;
    ws = new WebSocket("{{.}}");
    
    ws.onopen = function() {
        status.textContent = "Connected";
        status.className = "connected";
        log("✅ WebSocket connected", 'success');
    };
    
    ws.onclose = function() {
        status.textContent = "Disconnected";
        status.className = "disconnected";
        log("🔌 WebSocket connection closed", 'warning');
        ws = null;
    };
    
    ws.onmessage = function(evt) {
        log("📨 Received: " + evt.data, 'info');
    };
    
    ws.onerror = function(evt) {
        log("❌ Error: " + evt.data, 'error');
    };
};

document.getElementById("disconnect").onclick = function() {
    if (ws) {
        ws.close();
    }
};

document.getElementById("send").onclick = function() {
    if (!ws) {
        log("❌ Please connect WebSocket first", 'error');
        return;
    }
    var message = messageInput.value;
    ws.send(message);
    log("📤 Sent: " + message, 'success');
};

document.getElementById("sendTime").onclick = function() {
    if (!ws) return;
    ws.send("time");
    log("📤 Sent: time", 'success');
};

document.getElementById("sendStats").onclick = function() {
    if (!ws) return;
    ws.send("stats");
    log("📤 Sent: stats", 'success');
};

document.getElementById("sendPing").onclick = function() {
    if (!ws) return;
    ws.send("ping");
    log("📤 Sent: ping", 'success');
};

document.getElementById("clear").onclick = function() {
    output.innerHTML = '';
};

messageInput.addEventListener("keypress", function(e) {
    if (e.key === "Enter") {
        document.getElementById("send").click();
    }
});

// Auto connect
document.getElementById("connect").click();
</script>
</body>
</html>
`))