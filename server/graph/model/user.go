package model

import "time"

type User struct {
	ID                int       `json:"id"`
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	Password          string    `json:"password"`
	FollowingRoadmaps []int     `json:"followingRoadmaps"`
	CreatedRoadmaps   []int     `json:"createdRoadmaps"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}
