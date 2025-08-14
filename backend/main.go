package main

import (
    "log"
    "os"
    "voting-system/config"
    "voting-system/handlers"
    "voting-system/middleware"
    
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func main() {
    // Load environment variables
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    // Connect to database
    config.ConnectDB()
    defer config.DB.Close()

    // Setup Gin router
    r := gin.Default()

    // CORS middleware
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Content-Type", "Authorization"},
        AllowCredentials: true,
    }))

    // Routes
    api := r.Group("/api")
    {
        // Auth routes
        api.POST("/register", handlers.Register)
        api.POST("/login", handlers.Login)

        // Protected routes
        protected := api.Group("/")
        protected.Use(middleware.AuthMiddleware())
        {
            // Poll routes
            protected.POST("/polls", handlers.CreatePoll)
            protected.GET("/polls", handlers.GetPolls)
            protected.GET("/polls/:id", handlers.GetPollByID)
            protected.DELETE("/polls/:id", handlers.DeletePoll)
            
            // Vote routes
            protected.POST("/vote", handlers.Vote)
            protected.GET("/polls/:id/results", handlers.GetPollResults)
        }

        // Public routes (untuk polling tanpa auth)
        api.GET("/public/polls/:id", handlers.GetPollByID)
        api.POST("/public/vote", handlers.VotePublic)
        api.GET("/public/polls/:id/results", handlers.GetPollResults)
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Server running on port %s", port)
    r.Run(":" + port)
}