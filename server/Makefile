.PHONY: create up down force generate dataloader air

create:
	migrate create -ext sql -dir _db/migrations/ -seq add_checkpoints

up:
	migrate -database postgres://postgres:postgres@localhost/roadmap?sslmode=disable -path _db/migrations/ up
	
down:
	migrate -database postgres://postgres:postgres@localhost/roadmap?sslmode=disable -path _db/migrations/ down

force:
	migrate -database postgres://postgres:postgres@localhost/roadmap?sslmode=disable -path _db/migrations/ force 1

generate:
	go run github.com/99designs/gqlgen generate

dataloader:
	go run github.com/vektah/dataloaden UserLoader int *github.com/ivanwang123/roadmap/server/graph/model.User

air:
	air -c .air.toml