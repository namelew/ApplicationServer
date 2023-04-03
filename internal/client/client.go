package client

import (
	"log"

	"github.com/namelew/application-server/package/messages"
	"github.com/namelew/application-server/package/network"
	"github.com/spf13/cobra"
)

type Client struct {
	ServerAdress string
	ServerPort   string
}

func (c *Client) Run() {
	cli := cobra.Command{}

	cli.AddCommand(
		c.send(),
		c.ping(),
	)

	cli.Execute()
}

func (c *Client) send() *cobra.Command {
	return &cobra.Command{
		Use:   "send [string | json] [target]",
		Short: "send a message to another client",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}

func (c *Client) ping() *cobra.Command {
	var adress, port string

	cmd := cobra.Command{
		Use:   "ping",
		Short: "test connection with server",
		Run: func(cmd *cobra.Command, args []string) {
			conn := network.Connection{
				Adress: adress,
				Port:   port,
			}

			b := make(chan *messages.Message, 1)
			e := make(chan error, 1)

			conn.Open()

			m := messages.Message{
				Id:     1,
				Action: messages.PING,
				Body:   nil,
			}

			conn.Send(m)

			conn.Receive(b, e)

			resp := <-b

			if err := <-e; err != nil {
				log.Println(err.Error())
				conn.Close()
				return
			}

			log.Println(resp)

			conn.Close()
		},
	}

	cmd.Flags().StringVar(&adress, "adress", c.ServerAdress, "overrite default adress for test")
	cmd.Flags().StringVar(&port, "port", c.ServerPort, "overrite default port for test")

	return &cmd
}
