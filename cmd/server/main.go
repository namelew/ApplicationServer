package main

import (
	"github.com/namelew/application-server/internal/server"
	"github.com/namelew/application-server/internal/server/databases"
	"github.com/namelew/application-server/package/envoriment"
)

func main() {
	s := server.Server{
		Adress:   "localhost",
		Port:     "3001",
		Protocol: "tcp",
	}

	envoriment.LoadFile("./.env")

	d := databases.Database{}

	d.Connect()
	d.Disconnect()

	s.Run()
}
