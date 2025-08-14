package models

import (
    "time"
)

type Vote struct {
    ID        int       `json:"id"`
    PollID    int       `json:"poll_id"`
    UserID    int       `json:"user_id"`
    OptionID  int       `json:"option_id"`
    IPAddress string    `json:"ip_address"`
    CreatedAt time.Time `json:"created_at"`
}

type VoteRequest struct {
    PollID   int `json:"poll_id" binding:"required"`
    OptionID int `json:"option_id" binding:"required"`
}