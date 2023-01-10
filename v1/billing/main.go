package billing

import (
	"github.com/adomate-ads/api/middleware/auth"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	r.POST("/billing", auth.NotGuest, CreateBilling)
	r.GET("/billing", auth.NotGuest, GetBillings)
	r.GET("/billing/company/:id", auth.NotGuest, GetBillingsForCompany)
	r.GET("/billing/:id", auth.NotGuest, GetBilling)
	// TODO - Just test this one patch request for now... we can add more later once we know this one works
	r.PATCH("/billing/:id", auth.NotGuest, UpdateBilling)
	r.DELETE("/billing/:id", auth.NotGuest, DeleteBilling)
}
