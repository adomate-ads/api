package company

import (
	"github.com/adomate-ads/api/middleware/auth"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	r.POST("/company", auth.NotGuest, auth.InGroup("super-admin"), CreateCompany)
	r.GET("/company", auth.NotGuest, auth.InGroup("super-admin"), GetCompanies)
	r.GET("/company/:id", auth.NotGuest, GetCompany)
	r.DELETE("/company/:id", auth.NotGuest, auth.InGroup("super-admin"), DeleteCompany)
}
