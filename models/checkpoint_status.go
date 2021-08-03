package models

type NewCheckpointStatus struct {
	UserID       int
	RoadmapID    int
	CheckpointID int
}

type DeleteCheckpointStatus struct {
	RoadmapID     int
	UserIDs       []int
	CheckpointIDs []int
}

type GetCheckpointStatus struct {
	userID       int
	roadmapID    int
	checkpointID int
}

type CreateCheckpointStatus struct {
	userID       int
	roadmapID    int
	checkpointID int
}

type DeleteManyCheckpointStatus struct {
	roadmapID     int
	userIDs       []int
	checkpointIDs []int
}
