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
	conn     net.Conn
}

func (c *Connection) Open() error {
	conn, err := net.Dial(c.Protocol, c.Adress+":"+c.Port)

	c.conn = conn

	return err
}

func (c *Connection) Close() error {
	return c.conn.Close()
}

func (c *Connection) Send(m messages.Message) error {
	b, err := m.Serialize()

	if err != nil {
		return err
	}

	_, err = c.conn.Write(b)

	return err
}

func (c *Connection) Receive(buffer chan *messages.Message, berr chan error) {
	b := make([]byte, 1024)

	n, err := bufio.NewReader(c.conn).Read(b)

	if err != nil {
		buffer <- nil
		berr <- err
	}

	m, err := messages.Build(b[:n])

	buffer <- m
	berr <- err
}
