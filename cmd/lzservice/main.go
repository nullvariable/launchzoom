package main

import (
	b64 "encoding/base64"
	"log"
	"net"
	"os"

	"github.com/nullvariable/launchzoom/pkg/util"

	"github.com/kardianos/service"
)

const SockAddr = "/tmp/launchzoom.sock"

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() {
	if err := os.RemoveAll(SockAddr); err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("unix", SockAddr)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer l.Close()

	for {
		// Accept new connections, dispatching them to echoServer
		// in a goroutine.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		go echoServer(conn)
	}
}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "LaunchZoomService",
		DisplayName: "LaunchZoomService",
		Description: "LaunchZoomService Listener",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}

func echoServer(c net.Conn) {
	log.Printf("Client connected [%s]", c.RemoteAddr().Network())

	for {
		buf := make([]byte, 1024)
		nr, err := c.Read(buf)
		if err != nil {
			break
		}

		data := buf[0:nr]
		println("raw: " + string(data))
		url, _ := b64.StdEncoding.DecodeString(string(data))
		go util.Launch(string(url))
		c.Close()
		if err != nil {
			log.Fatal("Write: ", err)
		}
	}
}
