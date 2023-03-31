package main

import "github.com/namelew/application-server/internal/client"

func main() {
	c := client.Client{
		Server: "localhost:3001",
	}

	c.Run()
}
