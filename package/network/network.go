package network

import (
	"bufio"
	"net"

	"github.com/namelew/application-server/package/messages"
)

type Connection struct {
	Adress   string
	Port     string
	Protocol string
	Conn     net.Conn
}

func (c *Connection) Open() error {
	Conn, err := net.Dial(c.Protocol, c.Adress+":"+c.Port)

	c.Conn = Conn

	return err
}

func (c *Connection) Close() error {
	return c.Conn.Close()
}

func (c *Connection) Send(m messages.Message) error {
	b, err := m.Serialize()

	if err != nil {
		return err
	}

	_, err = c.Conn.Write(b)

	return err
}

func (c *Connection) Receive(buffer chan *messages.Message, berr chan error) {
	b := make([]byte, 1024)

	n, err := bufio.NewReader(c.Conn).Read(b)

	if err != nil {
		buffer <- nil
		berr <- err
	}

	m, err := messages.Build(b[:n])

	buffer <- m
	berr <- err
}
