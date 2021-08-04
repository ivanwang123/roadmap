package models

type GetCheckpointStatus struct {
	UserID       int
	RoadmapID    int
	CheckpointID int
}

type CreateCheckpointStatus struct {
	UserID       int
	RoadmapID    int
	CheckpointID int
}

type DeleteManyCheckpointStatus struct {
	RoadmapID     int
	UserIDs       []int
	CheckpointIDs []int
}
