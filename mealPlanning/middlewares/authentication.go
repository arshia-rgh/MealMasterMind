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
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	userID := getCurrentUser(token, context)

	context.Set("user", userID)
	context.Next()
}

func getCurrentUser(token string, context *gin.Context) int64 {
	req, err := http.NewRequest("GET", "http://localhost:8000/api/protected/me/", nil)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "err": err.Error()})
		return 0
	}

	req.Header.Set("Authorization", token)

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil || res.StatusCode != http.StatusOK {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "err": err})
		return 0
	}

	body, err := io.ReadAll(res.Body)
	if err != nil || len(body) == 0 {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "err": err})
		return 0
	}

	responseData, err := responseBodyToString(body)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "err": err.Error()})
	}

	userID, ok := responseData["id"].(int64)
	if !ok {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return 0
	}
	return userID
}

func responseBodyToString(body []byte) (map[string]interface{}, error) {
	var responseData map[string]interface{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		return nil, err
	}
	return responseData, nil
}
