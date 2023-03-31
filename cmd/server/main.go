package main

import (
	"github.com/namelew/application-server/internal/server"
	"github.com/namelew/application-server/internal/server/databases"
	"github.com/namelew/application-server/package/envoriment"
)

func main() {
	s := server.Server{
		Adress:   envoriment.GetVar("SRVADRESS"),
		Port:     envoriment.GetVar("SRVPORT"),
		Protocol: "tcp",
	}

	envoriment.LoadFile("./.env")

	d := databases.Database{}

	d.Connect()
	d.Migrate()
	d.Disconnect()

	s.Run()
}
