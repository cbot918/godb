package handler

import (
	"godb/src/godb"
	u "godb/util"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	db *godb.GODB
}

func New(db *godb.GODB) *Handler {

	h := new(Handler)

	h.db = db

	return h
}

func (h *Handler) Ping(c *gin.Context) {

	u.Logg(h.db.Ping())

}
