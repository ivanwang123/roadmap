package model

import "time"

type Roadmap struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatorID   int       `json:"creator"`
	Checkpoints []int     `json:"checkpoints"`
	Followers   []int     `json:"followers"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
