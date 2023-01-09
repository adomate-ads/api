package industry

import (
	"github.com/adomate-ads/api/middleware/auth"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	r.POST("/industry", auth.NotGuest, CreateIndustry)
	r.GET("/industry", auth.NotGuest, GetIndustries)
	r.GET("/industry/:industry", auth.NotGuest, GetIndustry)
	r.DELETE("/industry/:id", auth.NotGuest, DeleteIndustry)
}
