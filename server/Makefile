.PHONY: create up down force generate dataloader air

create:
	migrate create -ext sql -dir db/migrations/ -seq add_checkpoints

up:
	migrate -database postgres://postgres:postgres@localhost/roadmap?sslmode=disable -path db/migrations/ up
	
down:
	migrate -database postgres://postgres:postgres@localhost/roadmap?sslmode=disable -path db/migrations/ down

force:
	migrate -database postgres://postgres:postgres@localhost/roadmap?sslmode=disable -path db/migrations/ force 3

generate:
	go run github.com/99designs/gqlgen generate

dataloader:
	go run github.com/vektah/dataloaden RoadmapFollowerLoader int *github.com/ivanwang123/roadmap/server/graph/model.RoadmapFollower

air:
	air -c .air.toml