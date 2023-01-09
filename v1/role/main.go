package role

import (
	"github.com/adomate-ads/api/middleware/auth"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	r.POST("/role", auth.NotGuest, CreateRole)
	r.GET("/role", auth.NotGuest, GetRoles)
	r.GET("/role/:role", auth.NotGuest, GetRole)
	r.DELETE("/role/:id", auth.NotGuest, DeleteRole)
}
