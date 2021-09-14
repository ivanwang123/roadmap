.PHONY: create up down force generate dataloader air mock install-web build-web start-web

create:
	migrate create -ext sql -dir database/migrations/ -seq add_checkpoints

up:
	migrate -database postgres://postgres:postgres@localhost/roadmap?sslmode=disable -path database/migrations/ up
	
down:
	migrate -database postgres://postgres:postgres@localhost/roadmap?sslmode=disable -path database/migrations/ down

force:
	migrate -database postgres://postgres:postgres@localhost/roadmap?sslmode=disable -path database/migrations/ force 3

generate:
	go run github.com/99designs/gqlgen generate

dataloader:
	go run github.com/vektah/dataloaden RoadmapFollowerLoader int *github.com/ivanwang123/roadmap/server/graph/model.RoadmapFollower

air:
	air -c .air.toml

mock:
	mockery --name=Repository --recursive --dir=internal/user --output=./tests/mocks --filename=user_repo_mock.go --structname=UserMockRepo

install-web:
	cd web && npm run install
	
build-web: 
	cd web && npm run build

start-web: 
	cd web && npm start