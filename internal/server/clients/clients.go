package clients

import (
	"database/sql"
	"errors"
	"strconv"
	"time"
)

type Client struct {
	ID        uint64
	LastLogin time.Time
	CreatedAt time.Time
}

func (c *Client) Add(d *sql.DB) error {
	_, err := d.Exec("insert into clients(last_login,created_at) values ($1, $2)", c.LastLogin, c.CreatedAt)

	if err != nil {
		return errors.New("Unable to create new client. " + err.Error())
	}

	return nil
}

func (c *Client) Update(d *sql.DB) error {
	ret, err := d.Exec("update clients set last_login=$2 where id=$1", c.ID, c.CreatedAt)

	if err != nil {
		return errors.New("unable to update client data. " + err.Error())
	}

	rows, err := ret.RowsAffected()

	if err != nil || rows == 0 {
		return errors.New("unable to update client data. register not found")
	}

	return nil
}

func (c *Client) Remove(d *sql.DB) error {
	ret, err := d.Exec("delete from clients where id=$1", c.ID)

	if err != nil {
		return errors.New("unable to delete client data. " + err.Error())
	}

	rows, err := ret.RowsAffected()

	if err != nil || rows == 0 {
		return errors.New("unable to delete client data. register not found")
	}

	return nil
}

func (c *Client) Get(d *sql.DB, key interface{}) error {
	err := d.QueryRow("select id,last_login,created_at from clients where id = $1", key).Scan(&c.ID, &c.LastLogin, &c.CreatedAt)

	if err != nil {
		return errors.New("unable to bind client data. " + err.Error())
	}

	if c.ID == 0 {
		id := key.(int64)
		return errors.New("unable to find client " + strconv.FormatInt(id, 10) + " on database")
	}

	return nil
}
