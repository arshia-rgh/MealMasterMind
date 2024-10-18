package mails

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func RequestToSendMailGRPC(email, subject, link string) error {
	conn, err := grpc.NewClient("mailer-service:50001",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithConnectParams(grpc.ConnectParams{MinConnectTimeout: 5 * time.Second}),
	)
	if err != nil {
		return err
	}

	defer conn.Close()

	client := NewSendMailClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	responseResetLink, err := client.SendMail(ctx, &MailRequestResetLink{
		Email:   email,
		Subject: subject,
		Link:    link,
	},
	)
	if err != nil {
		return err
	}
	log.Println(responseResetLink.Result)
	return nil

}
