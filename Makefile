postgresinit:

postgres:
	docker exec -it // name of the psql ex. postgres16

createdb:
	docer exec -it postgress16 createdb --username=root --owner=root go-chat

dropdb:
	docer exec -it postgres16 dropdb go-chat

migrateup:
	migrate -path db/migrations -database "postgres://root:root@localhost:5432/go-chat?sslmode=disable" -verbose up"

migratedown:
	migrate -path db/migrations -database "postgres://root:root@localhost:5432/go-chat?sslmode=disable" -verbose down"

.PHONY: postgresinit postgres createdb dropdb migrateup migratedown
