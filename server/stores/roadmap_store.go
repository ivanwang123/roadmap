package stores

import (
	"fmt"

	"github.com/ivanwang123/roadmap/server/graph/model"
	"github.com/jmoiron/sqlx"
)

type RoadmapStore struct {
	*sqlx.DB
}

func (s *RoadmapStore) Create(input *model.NewRoadmap) (*model.Roadmap, error) {
	var roadmap model.Roadmap
	if err := s.Get(&roadmap, "INSERT INTO roadmaps (title, description, creator_id) VALUES ($1, $2, $3) RETURNING *",
		input.Title, input.Description, input.CreatorID); err != nil {
		return nil, err
	}
	return &roadmap, nil
}

func (s *RoadmapStore) GetAll() ([]*model.Roadmap, error) {
	roadmaps := []*model.Roadmap{}
	if err := s.Select(&roadmaps, "SELECT * FROM roadmaps"); err != nil {
		return []*model.Roadmap{}, err
	}
	fmt.Println("GET ALL ROADMAPS", roadmaps)
	return roadmaps, nil
}
