package main

import (
	"os"
	"github.com/ackdevelopment/accord/Discord"
	"fmt"
	"github.com/ackdevelopment/accord/Discord/GatewayRoutine"
)

func main() {
	defer os.Exit(0)
	/*var token string
	if lr, err := Discord.Login(Discord.LoginData{Email:"REDACTED", Password:"REDACTED"}); err == nil {
		fmt.Println(lr)
		if lr.MFA {
			var MFA int = 999999
			if mfa, merr := Discord.MFA(Discord.MFAData{Code:MFA, Ticket:lr.Ticket}); merr == nil {
				token = mfa.Token
			} else {
				fmt.Println(merr)
				return
			}
		}
	} else {
		fmt.Println(err)
		return
	}
	fmt.Println("Got token",token)*/

	var gateway string
	if ggr, err := Discord.GetGateway(); err == nil {
		gateway = ggr.Url + "?v=6&encoding=json"
	}
	fmt.Println("Connecting to gateway",gateway)

	ch := make(chan string)
	GatewayRoutine.GatewayRoutine(gateway, ch)
}