package httpService

import (
	"godb/src/godb"
	"godb/test/httpService/config"
	"log"

	"godb/test/httpService/route"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {

	// load config
	cfg, err := config.New("./test/configs", "app", "yaml")
	if err != nil {
		log.Fatalf("config.New error: %s", err)
	}

	// init db
	db := godb.New(
		cfg.Connections.Postgresql.DB,
		cfg.Connections.Postgresql.Host,
		cfg.Connections.Postgresql.Port,
		cfg.Connections.Postgresql.User,
		cfg.Connections.Postgresql.Password,
		cfg.Connections.Postgresql.MaxOpen,
		cfg.Connections.Postgresql.MaxIdle,
	)

	// init httpService
	httpEngine := gin.Default()
	httpService := route.New().RegistRoute(httpEngine, db)

	return httpService

}
