package course

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

type Psql struct {
	conn *pgx.Conn
}

func NewDB(conn string) *Psql {
	p := &Psql{}
	log.Print(conn)

	if p.ping() {
		log.Print("pgx ping ok")
	}

	if err := p.connect(conn); err != nil {
		log.Fatal("Can't connect to db \n", err)
	}

	if err := p.up(); err != nil {
		log.Fatal("Can't make migrations \n", err)
	}

	log.Print("pgx conn ok")

	return p
}

func (p *Psql) ping() bool {
	return true
}

func (p *Psql) up() error {
	migration, err := os.ReadFile("./migrations/0001_init.up.sql")
	log.Print(migration)

	return err
}

func (p *Psql) connect(conn string) error {
	client, err := pgx.Connect(context.Background(), conn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	if err != nil {
		panic(err)
	}
	p.conn = client

	return nil
}

func (p *Psql) Disconnect() error {
	log.Print("pgx disconn ok")
	return p.conn.Close(context.Background())
}
