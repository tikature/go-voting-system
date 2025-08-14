package models

import (
    "time"
)

type Poll struct {
    ID          int          `json:"id"`
    Title       string       `json:"title"`
    Description string       `json:"description"`
    CreatorID   int          `json:"creator_id"`
    ExpiresAt   *time.Time   `json:"expires_at"`
    IsActive    bool         `json:"is_active"`
    CreatedAt   time.Time    `json:"created_at"`
    Options     []PollOption `json:"options,omitempty"`
    TotalVotes  int          `json:"total_votes,omitempty"`
}

type PollOption struct {
    ID         int    `json:"id"`
    PollID     int    `json:"poll_id"`
    OptionText string `json:"option_text"`
    VoteCount  int    `json:"vote_count"`
}

type CreatePollRequest struct {
    Title       string     `json:"title" binding:"required"`
    Description string     `json:"description"`
    Options     []string   `json:"options" binding:"required,min=2"`
    ExpiresAt   *time.Time `json:"expires_at"`
}