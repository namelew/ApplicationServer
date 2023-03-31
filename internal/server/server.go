package server

import (
	"bufio"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/namelew/application-server/package/messages"
	"github.com/namelew/application-server/package/network"
)

type Server struct {
	Adress   string
	Port     string
	Protocol string
}

func (s *Server) Run() {
	var end = false
	listener, err := net.Listen(s.Protocol, s.Adress+":"+s.Port)

	if err != nil {
		log.Panic("Unable to create connection. ", err.Error())
	}

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		end = true
		listener.Close()
	}()

	for {
		c, err := listener.Accept()

		if end {
			break
		}

		if err != nil {
			log.Panic(err.Error())
		}

		go func(c *network.Connection) {
			b := make([]byte, 1024)

			n, err := bufio.NewReader(c.Conn).Read(b)

			if err != nil {
				log.Println("Failed to read request. ", err.Error())
				return
			}

			m, err := messages.Build(b[:n])

			if err != nil {
				log.Println("Failed to bind request. ", err.Error())
				return
			}

			switch m.Action {
			case messages.LIST:
			case messages.RELAY:
			case messages.RESPONSE:
			case messages.LOGIN:
			case messages.REGISTER:
			default:
				log.Println("Unknown action")
			}
		}(&network.Connection{
			Conn: c,
		})
	}
}
