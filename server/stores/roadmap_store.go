package stores

import (
	"github.com/ivanwang123/roadmap/server/graph/model"
	"github.com/jmoiron/sqlx"
)

type RoadmapStore struct {
	*sqlx.DB
}

const paginationLimit = 10

func (s *RoadmapStore) Create(input *model.NewRoadmap) (*model.Roadmap, error) {
	var roadmap model.Roadmap
	if err := s.Get(&roadmap, "INSERT INTO roadmaps (title, description, creator_id) VALUES ($1, $2, $3) RETURNING *",
		input.Title, input.Description, input.CreatorID); err != nil {
		return nil, err
	}
	return &roadmap, nil
}

func (s *RoadmapStore) GetByNewest(cursorID int, cursorValue string) ([]*model.Roadmap, error) {
	roadmaps := []*model.Roadmap{}
	if err := s.Select(&roadmaps, "SELECT * FROM roadmaps WHERE (created_at, id) <= ($1, $2) ORDER BY created_at DESC, id DESC LIMIT $3", cursorValue, cursorID, paginationLimit); err != nil {
		return nil, err
	}
	return roadmaps, nil
}

func (s *RoadmapStore) GetByOldest(cursorID int, cursorValue string) ([]*model.Roadmap, error) {
	roadmaps := []*model.Roadmap{}
	if err := s.Select(&roadmaps, "SELECT * FROM roadmaps WHERE (created_at, id) >= ($1, $2) ORDER BY created_at ASC, id ASC LIMIT $3", cursorValue, cursorID, paginationLimit); err != nil {
		return nil, err
	}
	return roadmaps, nil
}

func (s *RoadmapStore) GetByMostFollowers(cursorID int, cursorValue string) ([]*model.Roadmap, error) {
	roadmaps := []*model.Roadmap{}
	if err := s.Select(&roadmaps, "SELECT r.* FROM roadmaps r LEFT JOIN roadmap_followers rf ON rf.roadmap_id = r.id GROUP BY r.id HAVING (count(rf.roadmap_id), r.id) <= ($1, $2) ORDER BY count(rf.roadmap_id) DESC, r.id DESC LIMIT $3", cursorValue, cursorID, paginationLimit); err != nil {
		return nil, err
	}
	return roadmaps, nil
}

func (s *RoadmapStore) GetByMostCheckpoints(cursorID int, cursorValue string) ([]*model.Roadmap, error) {
	roadmaps := []*model.Roadmap{}
	if err := s.Select(&roadmaps, "SELECT r.* FROM roadmaps r LEFT JOIN checkpoints c ON c.roadmap_id = r.id GROUP BY r.id HAVING (count(c.id), r.id) <= ($1, $2) ORDER BY count(c.id) DESC, r.id DESC LIMIT $3", cursorValue, cursorID, paginationLimit); err != nil {
		return nil, err
	}
	return roadmaps, nil
}

func (s *RoadmapStore) GetByLeastCheckpoints(cursorID int, cursorValue string) ([]*model.Roadmap, error) {
	roadmaps := []*model.Roadmap{}
	if err := s.Select(&roadmaps, "SELECT r.* FROM roadmaps r LEFT JOIN checkpoints c ON c.roadmap_id = r.id GROUP BY r.id HAVING (count(c.id), r.id) >= ($1, $2) ORDER BY count(c.id) ASC, r.id ASC LIMIT $3", cursorValue, cursorID, paginationLimit); err != nil {
		return nil, err
	}
	return roadmaps, nil
}

func (s *RoadmapStore) GetAll() ([]*model.Roadmap, error) {
	roadmaps := []*model.Roadmap{}
	if err := s.Select(&roadmaps, "SELECT * FROM roadmaps"); err != nil {
		return nil, err
	}
	return roadmaps, nil
}

func (s *RoadmapStore) GetById(id int) (*model.Roadmap, error) {
	var roadmap model.Roadmap
	if err := s.Get(&roadmap, "SELECT * FROM roadmaps WHERE id = $1 LIMIT 1", id); err != nil {
		return nil, err
	}
	return &roadmap, nil
}

func (s *RoadmapStore) GetByCreatorId(creatorId int) ([]*model.Roadmap, error) {
	roadmaps := []*model.Roadmap{}
	if err := s.Select(&roadmaps, "SELECT * FROM roadmaps WHERE creator_id = $1", creatorId); err != nil {
		return nil, err
	}
	return roadmaps, nil
}
