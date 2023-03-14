package auth

import (
	"github.com/adomate-ads/api/models"
	"github.com/adomate-ads/api/pkg/auth"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"math"
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if companyID > math.MaxUint32 { // Add an upper bound check to ensure companyID can fit into a uint type.
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
		return
	}
	user := c.MustGet("x-user").(*models.User)
	if user.CompanyID != uint(companyID) && !auth.InGroup(user, "super-admin") {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You can only get information about your company"})
		return
	}
	c.Next()
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
