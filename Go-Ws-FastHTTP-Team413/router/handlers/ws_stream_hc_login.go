package handlers

import (
	"fmt"
	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
	"log"
	"strings"
)

//func (hd *HD) WsStreamHCLogin(w http.ResponseWriter, r *http.Request) {
//	auth := r.Header.Get("auth")
//	if auth != "a413|413" {
//		//===============================================================
//		w.WriteHeader(http.StatusUnauthorized)
//		_, err := w.Write([]byte("Client has been rejected | authentication fail"))
//		if err != nil {
//			return
//		}
//		return
//	}
//	conn, e1 := hd.WSU.Upgrade(w, r, nil)
//	if e1 != nil {
//		return
//	}
//	hd.WSCL.Lock()
//	fmt.Println("User a413 is connected!")
//	//fmt.Println(conn)
//	hd.WSCL.CM["a413"] = conn
//	hd.WSCL.Unlock()
//
//	// websocket.TextMessage = 1 | websocket.BinaryMessage = 2
//	//if err := conn.WriteMessage(websocket.TextMessage, []byte("1")); err != nil {
//	//	return
//	//}
//	//fmt.Fprint(w, "Goodbye world!")
//}

type WSC struct {
	L  bool
	WS *websocket.Conn
}

func (hd *HD) WsStreamHCLogin(ctx *fasthttp.RequestCtx) {
	auth := ctx.Request.Header.Peek("auth")
	addr := ctx.RemoteAddr()
	//ip := ctx.RemoteIP()
	fmt.Println(auth)
	u := "a413-" + addr.String()
	//===============================================================
	// Check if user is already logged in
	//===============================================================

	hd.WSCL.RLock()
	_, v := hd.WSCL.CM[u]
	hd.WSCL.RUnlock()
	if v == true {
		ctx.Error("Client ["+u+"] is already logged in", 503)
		return
	}
	// Add login ID - Window ID

	//===============================================================
	if strings.TrimSpace(string(auth)) != "a413|413" {
		//===============================================================
		//ctx.Response.WriteHeader(http.StatusUnauthorized)
		//_, err := w.Write([]byte("Client has been rejected | authentication fail"))
		//if err != nil {
		//	return
		//}
		ctx.Error("Client has been rejected | authentication fail", 503)
		return
	}
	fmt.Println("1")
	//=====================================================================
	err := hd.WSU.Upgrade(ctx, func(ws *websocket.Conn) {
		defer func(ws *websocket.Conn) {
			err := ws.Close()
			if err != nil {

			}
		}(ws)
		hd.WSCL.Lock()
		hd.WSCL.CM[u] = WSC{L: true, WS: ws}
		hd.WSCL.Unlock()
		fmt.Println("User [" + u + "] is connected!")
		for {
			mt, _, err := ws.ReadMessage()
			//fmt.Println(mt) // On 1006, message type is <-1>
			if mt == websocket.CloseMessage {
				hd.WSCL.Lock()
				_, e := hd.WSCL.CM[u]
				if e {
					hd.WSCL.CM[u] = WSC{L: false, WS: ws}
				}
				hd.WSCL.Unlock()
				log.Println("Client [" + u + "] closed connection | Message: CloseMessage")
				break
			}
			if mt == websocket.CloseAbnormalClosure {
				hd.WSCL.Lock()
				_, e := hd.WSCL.CM[u]
				if e {
					hd.WSCL.CM[u] = WSC{L: false, WS: ws}
				}
				hd.WSCL.Unlock()
				log.Println("Client [" + u + "] closed connection | Message: CloseAbnormalClosure")
				break
			}
			if err != nil {
				if websocket.IsCloseError(err, 1006) {
					hd.WSCL.Lock()
					_, e := hd.WSCL.CM[u]
					if e {
						hd.WSCL.CM[u] = WSC{L: false, WS: ws}
					}
					hd.WSCL.Unlock()
					log.Println("Connection [" + u + "] has been closed")
					break
				}
				//hd.WSCL.Lock()
				//delete(hd.WSCL.CM, "a413")
				//hd.WSCL.Unlock()
				//fmt.Println("User [" + u + "] has been deleted!")
				//break
			}
			//err = ws.WriteMessage(mt, []byte(""))
			//if err != nil {
			//	log.Println("write:", err)
			//	break
			//}
			//mt, message, err := ws.ReadMessage()
			//if err != nil {
			//	log.Println("read:", err)
			//	hd.WSCL.Lock()
			//	delete(hd.WSCL.CM, "a413")
			//	hd.WSCL.Unlock()
			//	break
			//}
			//log.Printf("recv: %s", message)
			//err = ws.WriteMessage(mt, message)
			//if err != nil {
			//	log.Println("write:", err)
			//	break
			//}
		}
	})
	//fmt.Println("2")

	if err != nil {
		fmt.Println("WS Handshake Error!")
		if _, ok := err.(websocket.HandshakeError); ok {
			log.Println(err)
		}
		return
	}
}

// https://github.com/fasthttp/websocket/blob/master/_examples/echo/client.go
// Cleanly close the connection by sending a close message and then
// waiting (with timeout) for the server to close the connection.

//			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
//			if err != nil {
//				log.Println("write close:", err)
//				return
//			}

/*
for {
            // Read message from browser
            msgType, msg, err := conn.ReadMessage()
            if err != nil {
                return
            }
			if mt == websocket.CloseMessage {
				disconnect <- true
			}
			if _, ok := err.(*websocket.CloseError); ok {
				disconnect <- true
			}
			if mt != websocket.TextMessage && mt != websocket.CloseMessage {
				w.logger.Println("warning: recieved unsupported message: ", mt, message)
				continue
			}


			//========================================================================
            // Print the message to the console
            fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

            // Write message back to browser
            if err = conn.WriteMessage(msgType, msg); err != nil {
                return
            }
        }
*/
