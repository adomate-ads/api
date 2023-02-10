package order

import (
	"github.com/adomate-ads/api/middleware/auth"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	r.POST("/order", auth.NotGuest, auth.InGroup("admin"), CreateOrder)
	r.GET("/order", auth.NotGuest, auth.InGroup("super-admin"), GetOrders)
	r.GET("/order/:id", auth.NotGuest, GetOrder)
	r.PATCH("/order/:id", auth.NotGuest, auth.InGroup("admin"), UpdateOrder)
	r.DELETE("/order/:id", auth.NotGuest, auth.InGroup("super-admin"), DeleteOrder)
}
