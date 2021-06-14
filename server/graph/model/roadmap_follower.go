package model

type RoadmapFollower struct {
	ID        int `json:"id"`
	UserID    int `json:"user"`
	RoadmapID int `json:"roadmap"`
}
