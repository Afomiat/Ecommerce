package main

import (
	"log"
	"time"

	"github.com/Afomiat/Ecommerce/config"
	"github.com/Afomiat/Ecommerce/delivery/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
    env := config.NewEnv()

    dbClient := config.ConnectMongoDB(env)
    db := dbClient.Database(env.DBName)

    timeout := time.Duration(env.ContextTimeout) * time.Second

    r := gin.Default()

    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"}, 
        AllowMethods:     []string{"GET", "POST", "PATCH","PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Content-Type", "Authorization"},
        AllowCredentials: true,  
        MaxAge:           12 * time.Hour,
    }))

    router.Setup(env, timeout, db, r)

    if err := r.Run(env.LocalServerPort); err != nil {
        log.Fatalf("Failed to start the server: %v", err)
    }
}
