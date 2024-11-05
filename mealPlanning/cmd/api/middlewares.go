package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"mealPlanning/grpc/user"

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
	if userID == 0 {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized: user not found or invalid token"})

	}

	context.Set("user", userID)
	context.Next()
}

func getCurrentUser(token string) (int64, error) {
	conn, err := grpc.NewClient("auth-service:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithConnectParams(grpc.ConnectParams{MinConnectTimeout: 10 * time.Second}),
	)

	defer conn.Close()

	if err != nil {
		return 0, err
	}

	log.Println("Connected to auth-service")

	client := user.NewAuthenticationClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.IsAuthenticated(ctx, &user.AuthReq{Token: token})
	if err != nil {
		return 0, err
	}
	log.Printf("recieved response from auth-service: %v", res.GetUserID())
	return res.UserID, nil
}

func RequestResponseLogger(c *gin.Context) {

}
