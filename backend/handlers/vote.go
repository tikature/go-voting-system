package handlers

import (
    "database/sql"
    "net/http"
    "strconv"
    "time"
    "voting-system/config"
    "voting-system/models"

    "github.com/gin-gonic/gin"
)

func Vote(c *gin.Context) {
    var req models.VoteRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID := c.GetInt("user_id")
    ipAddress := c.ClientIP()

    // Check if poll exists and is active
    var poll models.Poll
    err := config.DB.QueryRow(`SELECT id, expires_at, is_active FROM polls WHERE id = $1`, req.PollID).
        Scan(&poll.ID, &poll.ExpiresAt, &poll.IsActive)
    if err == sql.ErrNoRows {
        c.JSON(http.StatusNotFound, gin.H{"error": "Poll not found"})
        return
    }
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
        return
    }

    // Check if poll is active
    if !poll.IsActive {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Poll is not active"})
        return
    }

    // Check if poll has expired
    if poll.ExpiresAt != nil && time.Now().After(*poll.ExpiresAt) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Poll has expired"})
        return
    }

    // Check if user has already voted
    var existingVoteID int
    err = config.DB.QueryRow(`SELECT id FROM votes WHERE poll_id = $1 AND user_id = $2`, req.PollID, userID).
        Scan(&existingVoteID)
    if err == nil {
        c.JSON(http.StatusConflict, gin.H{"error": "You have already voted in this poll"})
        return
    }

    // Verify that the option belongs to this poll
    var optionPollID int
    err = config.DB.QueryRow(`SELECT poll_id FROM poll_options WHERE id = $1`, req.OptionID).
        Scan(&optionPollID)
    if err == sql.ErrNoRows {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid option"})
        return
    }
    if optionPollID != req.PollID {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Option does not belong to this poll"})
        return
    }

    // Start transaction
    tx, err := config.DB.Begin()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
        return
    }

    // Insert vote
    _, err = tx.Exec(`INSERT INTO votes (poll_id, user_id, option_id, ip_address) VALUES ($1, $2, $3, $4)`,
        req.PollID, userID, req.OptionID, ipAddress)
    if err != nil {
        tx.Rollback()
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record vote"})
        return
    }

    // Update vote count
    _, err = tx.Exec(`UPDATE poll_options SET vote_count = vote_count + 1 WHERE id = $1`, req.OptionID)
    if err != nil {
        tx.Rollback()
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update vote count"})
        return
    }

    // Commit transaction
    if err = tx.Commit(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit vote"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Vote recorded successfully"})
}

func VotePublic(c *gin.Context) {
    var req models.VoteRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ipAddress := c.ClientIP()

    // Check if poll exists and is active
    var poll models.Poll
    err := config.DB.QueryRow(`SELECT id, expires_at, is_active FROM polls WHERE id = $1`, req.PollID).
        Scan(&poll.ID, &poll.ExpiresAt, &poll.IsActive)
    if err == sql.ErrNoRows {
        c.JSON(http.StatusNotFound, gin.H{"error": "Poll not found"})
        return
    }
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
        return
    }

    // Check if poll is active
    if !poll.IsActive {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Poll is not active"})
        return
    }

    // Check if poll has expired
    if poll.ExpiresAt != nil && time.Now().After(*poll.ExpiresAt) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Poll has expired"})
        return
    }

    // Check if IP has already voted
    var existingVoteID int
    err = config.DB.QueryRow(`SELECT id FROM votes WHERE poll_id = $1 AND ip_address = $2`, req.PollID, ipAddress).
        Scan(&existingVoteID)
    if err == nil {
        c.JSON(http.StatusConflict, gin.H{"error": "This IP address has already voted in this poll"})
        return
    }

    // Verify that the option belongs to this poll
    var optionPollID int
    err = config.DB.QueryRow(`SELECT poll_id FROM poll_options WHERE id = $1`, req.OptionID).
        Scan(&optionPollID)
    if err == sql.ErrNoRows {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid option"})
        return
    }
    if optionPollID != req.PollID {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Option does not belong to this poll"})
        return
    }

    // Start transaction
    tx, err := config.DB.Begin()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
        return
    }

    // Insert vote (without user_id for public votes)
    _, err = tx.Exec(`INSERT INTO votes (poll_id, option_id, ip_address) VALUES ($1, $2, $3)`,
        req.PollID, req.OptionID, ipAddress)
    if err != nil {
        tx.Rollback()
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record vote"})
        return
    }

    // Update vote count
    _, err = tx.Exec(`UPDATE poll_options SET vote_count = vote_count + 1 WHERE id = $1`, req.OptionID)
    if err != nil {
        tx.Rollback()
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update vote count"})
        return
    }

    // Commit transaction
    if err = tx.Commit(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit vote"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Vote recorded successfully"})
}

func GetPollResults(c *gin.Context) {
    pollID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid poll ID"})
        return
    }

    poll, err := getPollWithOptions(pollID)
    if err == sql.ErrNoRows {
        c.JSON(http.StatusNotFound, gin.H{"error": "Poll not found"})
        return
    }
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch poll results"})
        return
    }

    c.JSON(http.StatusOK, poll)
}