package message

import (
	"sales_agent/common/log"

	"github.com/gin-gonic/gin"
)

func NewMessageHandler(c *gin.Context) {
	log.GetLogger().Info("new chat message: ", c.Request.Body)
}
