package middlewares

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	req, err := http.NewRequest("GET", "http://localhost:8000/api/protected/me/", nil)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	req.Header.Set("Authorization", token)

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil || res.StatusCode != http.StatusOK {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil || len(body) == 0 {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var responseData map[string]interface{}
	if err = json.Unmarshal(body, &responseData); err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID, ok := responseData["id"].(int64)
	if !ok {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	context.Set("user", userID)
	context.Next()
}
