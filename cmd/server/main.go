package main

import (
	"github.com/namelew/application-server/internal/server"
)

func main() {
	s := server.Server{
		Adress:   "localhost",
		Port:     "3001",
		Protocol: "tcp",
	}

	s.Run()
}
