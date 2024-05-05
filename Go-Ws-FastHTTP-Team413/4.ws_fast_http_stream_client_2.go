package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fasthttp/websocket"
)

func main() {
	pid := os.Getpid()
	fmt.Printf("My PID is %d\n", pid)

	h := http.Header{}
	h.Add("auth", "a413|413")
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8800/login", h) // ws://localhost:8090/ws
	if err != nil {
		if err == websocket.ErrBadHandshake {
			fmt.Println("Error on connecting to Websocket Server: Bad Handshake | Check your login info and be sure that user isn't already logged in")
			os.Exit(0)
		} else if websocket.IsUnexpectedCloseError(err) {

		} else {
			fmt.Println("Error connecting to Websocket Server | Error: ", err.Error())
			os.Exit(0)
		}
	}
	fmt.Println("Client has been connected!")
	defer func(c *websocket.Conn) {
		err = c.Close()
		if err != nil {

		}
	}(c)
	go WatchSystem("SW413 Tower Client", c)
	go func(c *websocket.Conn) {
		for {
			time.Sleep(500 * time.Millisecond)
			if err := c.WriteMessage(websocket.TextMessage, []byte("Architect413")); err != nil {
				return
			}
		}
	}(c)
	for {
		_, m, e2 := c.ReadMessage()
		//websocket.IsUnexpectedCloseError(e2,1005)
		if e2 != nil && websocket.IsCloseError(e2, 1006) {
			fmt.Println("Error on receiving message | Connection with server has been lost")
			break
		}
		fmt.Println(string(m))
	}
}

func WatchSystem(n string, c *websocket.Conn) {
	var s = make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
	v1 := <-s                       // interrupt or terminated
	if v1.String() == "interrupt" { //CTRL+C
		log.Println("<[" + n + "] IS SHUTDOWN> TERMINATED MANUALLY")
		err := c.Close()
		if err != nil {
			return
		}
		os.Exit(0)
	} else if v1.String() == "terminated" { //izbrisan process //SIGTERM
		log.Println("<["+n+"] IS SHUTDOWN>", "PROCCESS PID HAS BEEN DELETED")
		err := c.Close()
		if err != nil {
			return
		}
		os.Exit(0)
	} else if v1.String() == "stopped" {
		log.Println("<[" + n + "] IS WORKING InstanceName BACKGROUND>")
		err := c.Close()
		if err != nil {
			return
		}
		os.Exit(0)
	}
}

// https://github.com/fasthttp/websocket/blob/master/_examples/echo/client.go
// /usr/local/go-v1.8.6/go/bin/go build -o avalanchego-ade-v1.1 main.go