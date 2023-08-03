package resthandlers

import (
	"net/http"
	"time"

	"github.com/VinayakBagaria/go-cat-pictures/api/restutil"
	"github.com/gin-gonic/gin"
)

type ServerHandler interface {
	HealthCheck(*gin.Context)
}

type serverHandler struct {
	startAt time.Time
}

func NewServerHandler() ServerHandler {
	return &serverHandler{startAt: time.Now().UTC()}
}

func (h *serverHandler) HealthCheck(c *gin.Context) {
	now := time.Now().UTC()

	uptime := now.Sub(h.startAt)

	restutil.WriteAsJson(c, http.StatusOK, gin.H{
		"started_at": h.startAt.String(),
		"uptime":     uptime.String(),
		"ip_address": c.ClientIP(),
	})
}
