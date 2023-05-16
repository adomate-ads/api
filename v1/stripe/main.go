package stripe_v1

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	r.POST("/stripe/webhook", handleWebhook)
}
