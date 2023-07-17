package cloudflare

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

type SiteVerifyRequest struct {
	Token string `json:"cf_token"`
}

type SiteVerifyResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

func Verify(c *gin.Context) {
	var request SiteVerifyRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req, err := http.NewRequest(http.MethodPost, "https://challenges.cloudflare.com/turnstile/v0/siteverify", nil)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Add necessary request parameters.
	q := req.URL.Query()
	q.Add("secret", os.Getenv("CF_TURNSTILE_SECRET"))
	q.Add("response", request.Token)
	req.URL.RawQuery = q.Encode()

	// Make request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	var body SiteVerifyResponse
	if err = json.NewDecoder(resp.Body).Decode(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !body.Success {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": body.ErrorCodes})
		return
	}

	c.Next()
}
