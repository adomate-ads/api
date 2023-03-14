package auth

import (
	"github.com/adomate-ads/api/models"
	"github.com/adomate-ads/api/pkg/auth"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Auth(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get("user-id")
	if userId == nil {
		c.Set("x-guest", true)
		c.Next()
		return
	}

	user, err := models.GetUser(userId.(uint))
	if err == nil {
		c.Set("x-guest", false)
		c.Set("x-id", userId.(uint))
		c.Set("x-user", user)
		c.Set("x-auth-type", "cookie")
		c.Next()
		return
	}

	// If we get here, they had a cookie with an invalid user
	// so delete it.
	session.Delete("user-id")
	c.Set("x-guest", true)
	c.Next()
}

func NotGuest(c *gin.Context) {
	if c.GetBool("x-guest") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	c.Next()
}

func HasRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("x-user").(*models.User)
		if auth.HasRoleList(user, roles) || auth.InGroup(user, "super-admin") {
			c.Next()
			return
		}
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
	}
}

func SameCompany(c *gin.Context) {
	id := c.Param("id")
	companyID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := c.MustGet("x-user").(*models.User)
	if user.CompanyID != uint(companyID) && !auth.InGroup(user, "super-admin") {
		//c.JSON(http.StatusForbidden, gin.H{"error": "You can only get information about your company"})
		return
	}
	c.Next()
	return
}

func InGroup(group ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("x-user").(*models.User)
		for _, g := range group {
			if auth.InGroup(user, g) || auth.InGroup(user, "super-admin") {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
	}
}
