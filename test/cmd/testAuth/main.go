package main

import (
	"godb/test/service/auth/config"
	"godb/test/service/auth/rest"
	u "godb/util"
	"log"

	"godb/src/godb"
)

func main() {

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("new config error: %s", err)
	}
	u.Logg(cfg)

	db := godb.NewDB(
		cfg.Connections.Postgresql.DB,
		cfg.Connections.Postgresql.Host,
		cfg.Connections.Postgresql.Port,
		cfg.Connections.Postgresql.User,
		cfg.Connections.Postgresql.Password,
		cfg.Connections.Postgresql.MaxOpen,
		cfg.Connections.Postgresql.MaxIdle,
	)

	r := rest.NewRest(db)
	r.E.Run(":8201")
}
