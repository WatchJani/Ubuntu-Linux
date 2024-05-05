package handlers

import (
	"github.com/fasthttp/websocket"
	"sync"
)

type HD struct {
	Counter int
	WSU     websocket.FastHTTPUpgrader
	WSCL    *WSCL
}

func NewHD() *HD {
	return &HD{
		WSU: websocket.FastHTTPUpgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		WSCL: NewWSCL(),
	}
}

/*
var wsu = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
*/

type WSCL struct {
	CM map[string]WSC
	sync.RWMutex
}

func NewWSCL() *WSCL {
	return &WSCL{
		CM: map[string]WSC{},
	}
}
