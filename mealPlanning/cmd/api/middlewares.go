package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"mealPlanning/event"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"mealPlanning/grpc/user"

	"github.com/gin-gonic/gin"
)

func authentication(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	userID, userEmail, err := getCurrentUser(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "err": err.Error()})
	}
	if userID == 0 {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized: user not found or invalid token"})
	}
	if userEmail == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized: user not found or invalid token"})
	}

	context.Set("user", map[string]any{
		"id":    userID,
		"email": userEmail,
	})
	context.Next()
}

func getCurrentUser(token string) (int64, string, error) {
	conn, err := grpc.NewClient("auth-service:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithConnectParams(grpc.ConnectParams{MinConnectTimeout: 10 * time.Second}),
	)

	defer conn.Close()

	if err != nil {
		return 0, "", err
	}

	log.Println("Connected to auth-service")

	client := user.NewAuthenticationClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.IsAuthenticated(ctx, &user.AuthReq{Token: token})
	if err != nil {
		return 0, "", err
	}
	log.Printf("recieved response from auth-service: user(id, email) = (%v, %v)", res.GetUserID(), res.GetUserEmail())
	return res.UserID, res.UserEmail, nil
}

func requestResponseLogger(c *gin.Context) {
	var requestedUser string
	userInfo, exists := c.Get("user")

	if !exists {
		requestedUser = "Anonymous"
	} else {
		userEmail := userInfo.(map[string]any)["email"].(string)
		requestedUser = userEmail
	}

	// --Request--
	bodyBytes, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	log.Printf("Request: %s %s %s, user: %s", c.Request.Method, c.Request.URL, string(bodyBytes), requestedUser)

	go event.Publish("logs", map[string]string{
		"name":  "meal",
		"level": "info",
		"data":  fmt.Sprintf("Request: %s %s %s, user: %s", c.Request.Method, c.Request.URL, string(bodyBytes), requestedUser),
	})
	// --Response--
	responseBody := &bytes.Buffer{}
	writer := &responseWriter{body: responseBody, ResponseWriter: c.Writer}
	c.Writer = writer

	start := time.Now()
	c.Next()
	latency := time.Since(start)

	log.Printf("Response: %d %s %s", c.Writer.Status(), latency, responseBody.String())
	go event.Publish("logs", map[string]string{
		"name":  "meal",
		"level": "info",
		"data":  fmt.Sprintf("Response: %d %s %s", c.Writer.Status(), latency, responseBody.String()),
	})

}

type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
