package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// icon":{"iconType":"LIVE"}}
// "channelFeaturedContentRenderer":{"items":[{"vi
// <meta name="title" content=
func main() {
	channel := "https://www.youtube.com/channel/UCySNyY7ZlIyIeQ0T_8ah4Vg"
	channel = "https://www.youtube.com/watch?v=yIyqluvCvlo"
	req, err := http.NewRequest("GET", channel, nil)
	if err != nil {
		println(err)
	}
	resp, _ := http.DefaultClient.Do(req)
	content, _ := ioutil.ReadAll(resp.Body)
	f, _ := os.Create("title2.html")
	f.WriteString(string(content))
	roomOn := strings.Contains(string(content), `tWxhI3wYedQ`)
	// log.Println(string(content))
	println(roomOn)
}
