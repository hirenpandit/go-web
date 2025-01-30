package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func Handler() http.Handler {
	var upgrade = websocket.Upgrader{}
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		conn, err := upgrade.Upgrade(writer, request, nil)
		if err != nil {
			log.Printf("websocket upgrade failed: %v", err)
			return
		}
		defer func(conn *websocket.Conn) {
			err := conn.Close()
			if err != nil {
				log.Printf("websocket connection close failed: %v\n", err)
			}
		}(conn)
		for {
			mt, message, err := conn.ReadMessage()
			if err != nil {
				log.Printf("websocket read message failed: %v", err)
				break
			}
			log.Printf("websocket read message: %s of type %d \n", message, mt)
			err = conn.WriteMessage(1, []byte("ok"))
			if err != nil {
				log.Printf("websocket write message failed: %v", err)
			}
		}
	})
}
