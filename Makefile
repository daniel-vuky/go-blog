DB_URL=postgresql://root:secret@localhost:5432/go_blog?sslmode=disable
init_postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=root -d postgres:16-alpine
start_postgres:
	docker start postgres16
create_db:
	docker exec -it postgres16 createdb --username=root --owner=root go_blog
migrate_up:
	migrate -path database/migration -database "$(DB_URL)" -verbose up
migrate_down:
	migrate -path database/migration -database "$(DB_URL)" -verbose down
sqlc:
	sqlc generate
mock:
	mockgen -destination database/mock/store.go -package mockdb github.com/daniel-vuky/gogento-auth/database/sqlc Store
test:
	go test -v -cover -short ./...

.PHONY: init_postgres start_postgres stop_postgres create_db drop_db migrate_up migrate_down sqlc mock test
