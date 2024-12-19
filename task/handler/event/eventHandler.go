package event

import (
	"sales_agent/common/log"

	"github.com/gin-gonic/gin"
)

func NewEventHandler(c *gin.Context) {
	log.GetLogger().Info("event: ", c.Request.Body)
}
