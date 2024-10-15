package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type GateWayRequest struct {
	ServiceName string             `json:"service_name" binding:"required"`
	Auth        AuthServiceRequest `json:"auth_service_request,omitempty"`
}

type AuthServiceRequest struct {
	Action      string `json:"action" binding:"required"`
	UserID      int    `json:"user_id,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Username    string `json:"username,omitempty"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

func baseGateway(context *gin.Context) {
	var gateWay GateWayRequest

	err := context.ShouldBindJSON(gateWay)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

	switch gateWay.ServiceName {
	case "auth":
		authGateway(context, gateWay.Auth)
	default:
		context.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "service unknown"})
	}
}

func authGateway(context *gin.Context, auth AuthServiceRequest) {
	// TODO should publish message for all of them
	switch auth.Action {
	case "login":
		// TODO
	case "register":
		//TODO
	case "forgot-password", "reset-password", "request-reset-password":
		//TODO
	case "change_password":
		// TODO
	case "delete_user", "delete":
		// TODO
	case "update_user", "update":
		// TODO
	default:
		context.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "unknown action"})

	}
}
