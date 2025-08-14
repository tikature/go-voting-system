package handlers

import (
    "database/sql"
    "net/http"
    "voting-system/config"
    "voting-system/models"
    "voting-system/utils"

    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
    var req models.RegisterRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }

    // Insert user
    var userID int
    query := `INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id`
    err = config.DB.QueryRow(query, req.Username, req.Email, string(hashedPassword)).Scan(&userID)
    if err != nil {
        c.JSON(http.StatusConflict, gin.H{"error": "Username or email already exists"})
        return
    }

    // Generate JWT
    token, err := utils.GenerateJWT(userID, req.Username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "User registered successfully",
        "token":   token,
        "user": gin.H{
            "id":       userID,
            "username": req.Username,
            "email":    req.Email,
        },
    })
}

func Login(c *gin.Context) {
    var req models.LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    query := `SELECT id, username, email, password_hash FROM users WHERE username = $1`
    err := config.DB.QueryRow(query, req.Username).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)
    if err == sql.ErrNoRows {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // Verify password
    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // Generate JWT
    token, err := utils.GenerateJWT(user.ID, user.Username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Login successful",
        "token":   token,
        "user": gin.H{
            "id":       user.ID,
            "username": user.Username,
            "email":    user.Email,
        },
    })
}