DB_URL=postgresql://root:secret@localhost:2345/seatmap?sslmode=disable
# DB_SOURCE=postgresql://root:secret@postgres15seatmap:2345/seatmap?sslmode=disable
docker_prepare: 
	docker pull postgres:15.2-alpine
	docker network create seatmap-network
	
postgres:
	docker run --name postgres15seatmap --network seatmap-network -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 2345:5432 -d postgres:15.2-alpine

createdb:
	docker exec -it postgres15seatmap createdb --username=root --owner=root seatmap

dropdb:
	docker exec -it postgres15seatmap dropdb seatmap

server: 
	go run main.go

docker_create_network: 
	docker network create seatmap-network

docker_build:
	docker build -t seatmapbackend:latest .

docker_run:
	docker run --name seatmapbackend --network seatmap-network -e DB_SOURCE="postgresql://root:secret@postgres15seatmap:2345/seatmap?sslmode=disable" -p 8080:8080 seatmapbackend:latest

docker_clean: 
	docker stop postgres15seatmap
	docker rm postgres15seatmap
	docker rmi postgres:15.2-alpine
	docker stop seatmapbackend
	docker rm seatmapbackend
	docker rmi seatmapbackend
	docker network rm seatmap-network

# The following commands require golang-migrate CLI. https://github.com/golang-migrate/migrate and run on local
migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

new_migration: 
# make name=your_migration_name new_migration
	migrate create -ext sql -dir db/migration -seq $(name)

.PHONY: pull_docker_img postgres createdb dropdb docker_clean migrateup migrateup1 migratedown migratedown1 new_migration server
