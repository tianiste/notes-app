package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"notes-app/internal/db"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func main() {
	ctx := context.Background()
	if err := db.Init(ctx); err != nil {
		log.Fatal(err)
	}
	client := db.Client
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("connected")
	defer db.Close(ctx)
	router := gin.Default()
	api := router.Group("/api")
	api.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
}
