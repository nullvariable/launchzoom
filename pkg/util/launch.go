package util

import (
	"fmt"
	"net/url"
	"os/exec"
	"strings"
)

func Launch(arg string) {
	url, err := url.Parse(arg)
	if err != nil {
		panic(err)
	}
	query := url.Query()
	path := strings.Split(url.Path, "/")
	fmt.Println(path)
	meetingid := path[len(path)-1]
	zoommtg := fmt.Sprintf("zoommtg://%s/join?confno=%s", url.Host, meetingid)

	if query.Get("pwd") != "" {
		zoommtg += fmt.Sprintf("&pwd=%s", query.Get("pwd"))
	}
	fmt.Println("final: " + zoommtg)
	cmd := exec.Command("xdg-open", zoommtg)
	cmd.Start()
}
