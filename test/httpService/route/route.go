package route

import (
	"godb/src/godb"
	"godb/test/httpService/handler"

	"github.com/gin-gonic/gin"
)

type Route struct {
}

func New() *Route {
	r := new(Route)
	return r
}

func (r *Route) RegistRoute(engine *gin.Engine, godb *godb.GODB) *gin.Engine {

	h := handler.New(godb)
	engine.GET("/ping", h.Ping)

	return engine
}
