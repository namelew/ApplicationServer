package main

import (
	"github.com/namelew/application-server/internal/client"
	"github.com/namelew/application-server/package/envoriment"
)

func main() {
	c := client.Client{
		Server: envoriment.GetVar("SRVADRESS") + ":" + envoriment.GetVar("SRVPORT"),
	}

	c.Run()
}
