package preregistration

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	r.POST("/preregistration", CreatePreRegistration)
}
