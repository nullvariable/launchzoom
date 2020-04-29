package util

import (
	b64 "encoding/base64"
	"log"
	"net"
	"os"
)

const SockAddr = "/tmp/launchzoom.sock"

func WriteToSock(arg string) error {

	c, err := net.Dial("unix", SockAddr)
	sEnc := b64.StdEncoding.EncodeToString([]byte(os.Args[1]))
	println("Sending: " + string(sEnc))
	_, err = c.Write([]byte(sEnc))
	if err != nil {
		log.Fatal("write error:", err)
	}
	defer c.Close()
	return err
}
