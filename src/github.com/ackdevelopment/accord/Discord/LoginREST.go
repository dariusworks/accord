package Discord

import (
	"encoding/json"
	"net/http"
	"time"
	"strings"
	"io/ioutil"
)

var client = http.Client{Timeout:time.Second*5}
var transport = http.Transport{}
var baseAddress = "https://discordapp.com/api/v6/"

type LoginData struct {
	Email string	`json:"email"`
	Password string	`json:"password"`
}

type LoginResponse struct {
	MFA bool
	Ticket string
	Token string
}

func Login(LData LoginData) (lr LoginResponse, err error) {
	d,_ := json.Marshal(LData)
	req, _ := http.NewRequest("POST", baseAddress + "auth/login", strings.NewReader(string(d)))
	req.Header.Set("Content-Type","application/json")
	if r, rerr := client.Do(req); rerr == nil {
		if rd, cerr := ioutil.ReadAll(r.Body); cerr == nil {
			json.Unmarshal(rd,&lr)
			return lr, nil
		} else {
			return LoginResponse{}, cerr
		}
	} else {
		return LoginResponse{}, rerr
	}
}

type MFAData struct {
	Code int `json:"code"`
	Ticket string `json:"ticket"`
}

type MFAResponse struct {
	Token string
}

func MFA(data MFAData) (mr MFAResponse, err error) {
	d,_ := json.Marshal(data)
	req, _ := http.NewRequest("POST", baseAddress + "auth/mfa/totp", strings.NewReader(string(d)))
	req.Header.Set("Content-Type", "application/json")
	if r, rerr := client.Do(req); rerr == nil {
		if rd, cerr := ioutil.ReadAll(r.Body); cerr == nil {
			json.Unmarshal(rd,&mr)
			return mr,nil
		} else {
			return MFAResponse{}, cerr
		}
	} else {
		return MFAResponse{}, rerr
	}
}