package GatewayRoutine

type GatewayPayload struct {
	OP int `json:"op"`
	/*
	0  Dispatch 			 receive
	1  Heartbeat 			 send/receive
	2  identify 			 send
	3  status_update 		 send
	4  voice_state_update 	 send
	5  voice_server_ping 	 send
	6  resume 				 send
	7  reconnect 			 receive
	8  request_guild_members send
	9  invalid_session 		 receive
	10 hello 				 receive
	11 heartbeat_ack 		 receive
	*/
	Data map[string]interface{} `json:"d"`
	Sequence *int `json:"s"`
	T string `json:"t"`
}

type GatewayHello struct { //receive only
	HeartbeatInterval int 	`mapstructure:"heartbeat_interval"`
	Trace []string 			`mapstructure:"_trace"`
}

type GatewayHeartbeat struct { //send/receive
	OP int `mapstructure:"op"`
	D int  `mapstructure:"d"`
}

type GatewayHeartbeatACK struct { //receive only
	OP int `mapstructure:"op"`
}