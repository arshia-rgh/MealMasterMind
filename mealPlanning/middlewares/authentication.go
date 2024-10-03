package middlewares

import (
	"encoding/json"
	"errors"
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

	userID, err := getCurrentUser(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "err": err.Error()})
	}

	context.Set("user", userID)
	context.Next()
}

func getCurrentUser(token string) (int64, error) {
	req, err := http.NewRequest("GET", "http://localhost:8000/api/protected/me/", nil)
	if err != nil {
		return 0, err
	}

	req.Header.Set("Authorization", token)

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil || res.StatusCode != http.StatusOK {
		return 0, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil || len(body) == 0 {
		return 0, err
	}

	responseData, err := responseBodyToString(body)

	if err != nil {
		return 0, err
	}

	userID, ok := responseData["id"].(int64)
	if !ok {
		return 0, errors.New("user id is not int")
	}
	return userID, nil
}

func responseBodyToString(body []byte) (map[string]interface{}, error) {
	var responseData map[string]interface{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		return nil, err
	}
	return responseData, nil
}
