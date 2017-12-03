package Discord

import (
	"net/http"
	"strings"
	"io/ioutil"
	"encoding/json"
)

type GGResponse struct {
	Url string
}

func GetGateway() (GGR GGResponse, err error) {
	req, _ := http.NewRequest("GET", baseAddress + "gateway", strings.NewReader(""))
	if r, rerr := client.Do(req); err == nil {
		if d, cerr := ioutil.ReadAll(r.Body); cerr == nil {
			json.Unmarshal(d,&GGR)
			return GGR, nil
		} else {
			return GGR, cerr
		}
	} else {
		return GGR, rerr
	}
}