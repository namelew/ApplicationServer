package databases

import (
	"database/sql"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/namelew/application-server/package/envoriment"
)

type Query interface {
	Add(d *sql.DB) error
	Update(d *sql.DB) error
	Remove(d *sql.DB) error
	Get(d *sql.DB, key interface{}) error
}

type Database struct {
	db *sql.DB
}

func (d *Database) Connect() {
	db, err := sql.Open("sqlite3", envoriment.GetVar("DBPATH"))

	if err != nil {
		log.Panic("unable to connect to database. ", err.Error())
	}

	d.db = db
}

func (d *Database) Disconnect() {
	if err := d.db.Close(); err != nil {
		log.Fatal("Unable to close connection with database. ", err.Error())
	}
}

func (d *Database) Migrate() {
	data, err := os.ReadFile("./migrations/" + envoriment.GetVar("DBNAME") + ".up.sql")

	if err != nil {
		log.Panic("Unable to load migrate configs. ", err.Error())
	}

	sanitaze := func(s string) string {
		trash := []string{
			"\n", "\b", "\t", "\a", "\r", "\f", "\v",
		}

		for i := range trash {
			s = strings.ReplaceAll(s, trash[i], "")
		}

		return s
	}

	for _, command := range strings.Split(sanitaze(string(data)), ";") {
		_, err = d.db.Exec(command)

		if err != nil {
			log.Fatal("Unable to execute migration step ", command, ".", err.Error())
		}
	}
}

func (d *Database) Add(q Query) error {
	return q.Add(d.db)
}

func (d *Database) Update(q Query) error {
	return q.Update(d.db)
}

func (d *Database) Remove(q Query) error {
	return q.Remove(d.db)
}

func (d *Database) Get(q Query, key interface{}) error {
	return q.Get(d.db, key)
}
