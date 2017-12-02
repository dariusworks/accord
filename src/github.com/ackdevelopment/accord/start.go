package main

import (
	"os"
	"github.com/ackdevelopment/accord/Discord"
	"fmt"
)

func main() {
	defer os.Exit(0)
	if lr, err := Discord.Login(Discord.LoginData{"REDACTED", "REDACTED"}); err == nil {
		fmt.Println(lr)
		if lr.MFA {
			var MFA int = 999999
			if mfa, merr := Discord.MFA(Discord.MFAData{Code:MFA, Ticket:lr.Ticket}); merr == nil {
				fmt.Println(mfa)
			} else {
				fmt.Println(merr)
			}
		}
	} else {
		fmt.Println(err)
	}
}