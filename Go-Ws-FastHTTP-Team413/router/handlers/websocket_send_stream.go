package handlers

import (
	"fmt"
	"github.com/fasthttp/websocket"
)

func (hd *HD) WebsocketSendHC() {
	for {
		//fmt.Println("1111")
		//time.Sleep(200 * time.Millisecond)
		hd.Counter++
		//hd.WSCL.RLock()
		//c, ok := hd.WSCL.CM["a413"]
		//hd.WSCL.RUnlock()
		//if !ok {
		//	fmt.Println("Key not found")
		//	//fmt.Println("Key found value is: ", c)
		//
		//}
		//fmt.Println("2222")
		//if c != nil {
		//	fmt.Println("3333")
		//	if e1 := c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%v", hd.Counter))); e1 != nil {
		//		//_ = c.Close()
		//		fmt.Println("4444")
		//		hd.WSCL.Lock()
		//		delete(hd.WSCL.CM, "a413")
		//		hd.WSCL.Unlock()
		//	}
		//}
		//hd.WSCL.RLock()
		//c, ok := hd.WSCL.CM["a413"]
		//hd.WSCL.RUnlock()
		//if !ok {
		//	fmt.Println("Key not found")
		//	//fmt.Println("Key found value is: ", c)
		//
		//}
		hd.WSCL.Lock()
		for u, c := range hd.WSCL.CM {
			//fmt.Println("2222")
			if c.L {
				//fmt.Println("3333")
				//c.SetWriteDeadline(time.Second * 0)
				if e1 := c.WS.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%v", hd.Counter))); e1 != nil {
					//_ = c.Close()
					//fmt.Println("4444")
					delete(hd.WSCL.CM, u)
					fmt.Println("User [" + u + "] has been deleted!")

				}
				continue
			}
			delete(hd.WSCL.CM, u)
			fmt.Println("User [" + u + "] has been deleted!")
		}
		hd.WSCL.Unlock()
	}
	// websocket.TextMessage = 1 | websocket.BinaryMessage = 2
	//if err := conn.WriteMessage(websocket.TextMessage, []byte("1")); err != nil {
	//	return
	//}
	//fmt.Fprint(w, "Goodbye world!")
}

//func (hd *HD) echoView(ctx *fasthttp.RequestCtx) {
//	err := upgrader.Upgrade(ctx, func(ws *websocket.Conn) {
//		defer ws.Close()
//		for {
//			mt, message, err := ws.ReadMessage()
//			if err != nil {
//				log.Println("read:", err)
//				break
//			}
//			log.Printf("recv: %s", message)
//			err = ws.WriteMessage(mt, message)
//			if err != nil {
//				log.Println("write:", err)
//				break
//			}
//		}
//	})
//
//	if err != nil {
//		if _, ok := err.(websocket.HandshakeError); ok {
//			log.Println(err)
//		}
//		return
//	}
//}
