package model

import "time"

type Checkpoint struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Instructions string    `json:"instructions"`
	Links        []string  `json:"links"`
	IsCompleted  *bool     `json:"isCompleted"`
	RoadmapID    int       `json:"roadmap"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
