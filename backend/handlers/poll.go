package handlers

import (
    "database/sql"
    "net/http"
    "strconv"
    "voting-system/config"
    "voting-system/models"

    "github.com/gin-gonic/gin"
)

func CreatePoll(c *gin.Context) {
    var req models.CreatePollRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID := c.GetInt("user_id")

    // Insert poll
    var pollID int
    query := `INSERT INTO polls (title, description, creator_id, expires_at) VALUES ($1, $2, $3, $4) RETURNING id`
    err := config.DB.QueryRow(query, req.Title, req.Description, userID, req.ExpiresAt).Scan(&pollID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create poll"})
        return
    }

    // Insert options
    for _, optionText := range req.Options {
        _, err := config.DB.Exec(`INSERT INTO poll_options (poll_id, option_text) VALUES ($1, $2)`, pollID, optionText)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create poll options"})
            return
        }
    }

    // Fetch created poll with options
    poll, err := getPollWithOptions(pollID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch created poll"})
        return
    }

    c.JSON(http.StatusCreated, poll)
}

func GetPolls(c *gin.Context) {
    userID := c.GetInt("user_id")

    query := `
        SELECT p.id, p.title, p.description, p.creator_id, p.expires_at, p.is_active, p.created_at,
               COALESCE(SUM(po.vote_count), 0) as total_votes
        FROM polls p
        LEFT JOIN poll_options po ON p.id = po.poll_id
        WHERE p.creator_id = $1
        GROUP BY p.id, p.title, p.description, p.creator_id, p.expires_at, p.is_active, p.created_at
        ORDER BY p.created_at DESC
    `

    rows, err := config.DB.Query(query, userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch polls"})
        return
    }
    defer rows.Close()

    var polls []models.Poll
    for rows.Next() {
        var poll models.Poll
        err := rows.Scan(&poll.ID, &poll.Title, &poll.Description, &poll.CreatorID, 
                        &poll.ExpiresAt, &poll.IsActive, &poll.CreatedAt, &poll.TotalVotes)
        if err != nil {
            continue
        }
        polls = append(polls, poll)
    }

    c.JSON(http.StatusOK, polls)
}

func GetPollByID(c *gin.Context) {
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
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch poll"})
        return
    }

    c.JSON(http.StatusOK, poll)
}

func DeletePoll(c *gin.Context) {
    pollID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid poll ID"})
        return
    }

    userID := c.GetInt("user_id")

    // Check if user owns the poll
    var creatorID int
    err = config.DB.QueryRow(`SELECT creator_id FROM polls WHERE id = $1`, pollID).Scan(&creatorID)
    if err == sql.ErrNoRows {
        c.JSON(http.StatusNotFound, gin.H{"error": "Poll not found"})
        return
    }
    if creatorID != userID {
        c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your own polls"})
        return
    }

    // Delete poll (cascade will delete options and votes)
    _, err = config.DB.Exec(`DELETE FROM polls WHERE id = $1`, pollID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete poll"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Poll deleted successfully"})
}

func getPollWithOptions(pollID int) (*models.Poll, error) {
    var poll models.Poll
    query := `SELECT id, title, description, creator_id, expires_at, is_active, created_at FROM polls WHERE id = $1`
    err := config.DB.QueryRow(query, pollID).Scan(&poll.ID, &poll.Title, &poll.Description, 
                                                  &poll.CreatorID, &poll.ExpiresAt, &poll.IsActive, &poll.CreatedAt)
    if err != nil {
        return nil, err
    }

    // Get options
    optionsQuery := `SELECT id, poll_id, option_text, vote_count FROM poll_options WHERE poll_id = $1 ORDER BY id`
    rows, err := config.DB.Query(optionsQuery, pollID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var options []models.PollOption
    var totalVotes int
    for rows.Next() {
        var option models.PollOption
        err := rows.Scan(&option.ID, &option.PollID, &option.OptionText, &option.VoteCount)
        if err != nil {
            continue
        }
        options = append(options, option)
        totalVotes += option.VoteCount
    }

    poll.Options = options
    poll.TotalVotes = totalVotes

    return &poll, nil
}