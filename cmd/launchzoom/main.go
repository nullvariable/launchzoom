package main

/**
 * Takes a zoom url, and kicks it to an active system service if there is one
 * Keeps zoom from being a child of other apps like Slack if launched there
 * Lets us replace our default browser with something that bypasses the whole
 * launch a browser to run xdg-open for zoomlinks
 */
// https://XXX.zoom.us/j/999999999
// https://XXX.zoom.us/j/99999999?pwd=alosghergeruhgihugkfjhsrkgjh

import (
	"fmt"
	"os"

	"github.com/nullvariable/launchzoom/pkg/util"
)

const SockAddr = "/tmp/launchzoom.sock"

func main() {
	if len(os.Args) > 1 && os.Args[1] != "" { // if we're passed a url, try to connect and then exit.
		fmt.Println(os.Args[1])
		err := util.WriteToSock(os.Args[1])
		if err != nil { // no socket, so just launch.
			fmt.Println("no socket found :(")
			// panic(err)
			util.Launch(os.Args[1])
			os.Exit(1)
		}
	} else { // no url
		fmt.Println("Please pass an argument")
	}
}

// func launch(arg string) {
// 	url, err := url.Parse(arg)
// 	if err != nil {
// 		panic(err)
// 	}
// 	query := url.Query()
// 	path := strings.Split(url.Path, "/")
// 	fmt.Println(path)
// 	meetingid := path[len(path)-1]
// 	zoommtg := fmt.Sprintf("zoommtg://%s/join?confno=%s", url.Host, meetingid)

// 	if query.Get("pwd") != "" {
// 		zoommtg += fmt.Sprintf("&pwd=%s", query.Get("pwd"))
// 	}
// 	fmt.Println("final: " + zoommtg)
// 	cmd := exec.Command("xdg-open", zoommtg)
// 	cmd.Start()
// }
