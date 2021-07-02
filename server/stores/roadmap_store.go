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

func (s *RoadmapStore) GetByNewest(cursor string) ([]*model.Roadmap, error) {
	roadmaps := []*model.Roadmap{}
	if err := s.Select(&roadmaps, "SELECT * FROM roadmaps WHERE created_at < $1 ORDER BY created_at DESC LIMIT $2", cursor, paginationLimit); err != nil {
		return nil, err
	}
	return roadmaps, nil
}

func (s *RoadmapStore) GetByOldest(cursor string) ([]*model.Roadmap, error) {
	return nil, nil
}

func (s *RoadmapStore) GetByMostFollowers(cursor string) ([]*model.Roadmap, error) {
	// TODO: Figure out how to format cursor for count & id
	// SELECT * FROM roadmaps LEFT JOIN roadmap_followers rf ON rf.roadmap_id = id GROUP BY id HAVING (count(rf.user_id), id) <= (1, 1) ORDER BY count(rf.user_id) DESC, id DESC LIMIT 30;
	return nil, nil
}

func (s *RoadmapStore) GetByMostCheckpoints(cursor string) ([]*model.Roadmap, error) {
	return nil, nil
}

func (s *RoadmapStore) GetByLeastCheckpoints(cursor string) ([]*model.Roadmap, error) {
	return nil, nil
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
