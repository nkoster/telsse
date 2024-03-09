package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

var (
	clients    = make(map[chan string]struct{})
	mu         sync.Mutex
	uiPath     string
	telnetPort string
	logging    bool
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	uiPath = os.Getenv("UI")
	httpPort := os.Getenv("HTTP_PORT")
	telnetPort = os.Getenv("TELNET_PORT")
	logging, err = strconv.ParseBool(os.Getenv("LOG"))
	if err != nil {
		log.Fatalf("Error converting LOG to boolean: %v", err)
	}

	http.Handle("/", http.FileServer(http.Dir(uiPath)))
	http.HandleFunc("/sse", sseHandler)

	go startTelnetServer()

	log.Printf("HTTP server listening on port %s", httpPort)
	log.Fatal(http.ListenAndServe(":"+httpPort, nil))
}

func startTelnetServer() {
	listener, err := net.Listen("tcp", ":"+telnetPort)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	log.Printf("Telnet server listening on port %s", telnetPort)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleTelnetConnection(conn)
	}
}

func handleTelnetConnection(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		message := scanner.Text()
		if logging {
			fmt.Println("Received via Telnet:", message)
		}
		sendMessageToClients(message)
	}
}

func sendMessageToClients(message string) {
	mu.Lock()
	defer mu.Unlock()
	for client := range clients {
		client <- message
	}
}

func sseHandler(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	client := make(chan string)
	mu.Lock()
	clients[client] = struct{}{}
	mu.Unlock()

	defer func() {
		mu.Lock()
		delete(clients, client)
		mu.Unlock()
	}()

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	for {
		msg, open := <-client
		if !open {
			break
		}

		fmt.Fprintf(w, "data: %s\n\n", msg)
		flusher.Flush()
	}
}
