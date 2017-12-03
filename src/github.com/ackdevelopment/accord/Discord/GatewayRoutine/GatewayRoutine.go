package GatewayRoutine

import (
	"github.com/gorilla/websocket"
	"log"
	"time"
	"encoding/json"
	"github.com/mitchellh/mapstructure"
)

var Gateway string
var GWConn websocket.Conn

func GatewayRoutine(GWURL string, ch chan string) {
	Gateway = GWURL

	c, _, err := websocket.DefaultDialer.Dial(Gateway,nil)
	if err != nil {
		log.Fatal("dial: ", err)
	}
	defer c.Close()

	heartbeat := time.Tick(0)
	payloads := make(chan GatewayPayload)

	go func() {
		defer c.Close()
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println(nil)
			}
			log.Println("recieved message containing\n", string(message))

			var GWPL GatewayPayload
			json.Unmarshal(message, &GWPL)
			log.Println(GWPL)
			payloads <- GWPL
		}
	}()

	var sequence *int = nil

	for {
		select {
			case <- heartbeat:
				log.Println("Sending heartbeat")
				if out, err := json.Marshal(GatewayPayload{OP:1, Data:map[string]interface{}{"op":1, "d":251}, Sequence:sequence}); err == nil {
					c.WriteMessage(1,out)
				}
			case pl := <- payloads:
				switch pl.OP {
				case 1: //Heartbeat
					//do nothing.
				case 10: //Hello
					var GWH GatewayHello
					err := mapstructure.Decode(pl.Data, &GWH)
					if err == nil {
						heartbeat = time.Tick(time.Millisecond * time.Duration(GWH.HeartbeatInterval))
					} else {
						log.Println(err)
					}
				case 11: //Heartbeat ACK
					//do nothing.
				}
				if pl.Sequence != nil {
					log.Println("")
					sequence = pl.Sequence
				}
		}
	}
}