GOOSE_DRIVER=mysql
GOOSE_DBSTRING="root:@tcp(localhost:3306)/user"
OS=linux

build:
	GOARCH=amd64 GOOS=linux go build -o ./bin/linux ./cmd/main.go
	GOARCH=amd64 GOOS=darwin go build -o ./bin/darwin ./cmd/main.go
	GOARCH=amd64 GOOS=windows go build -o ./bin/windows ./cmd/main.go

build-run:
	make build
	./bin/${OS}/main

run:
	go build -o ./bin/linux ./cmd/main.go
	./bin/${OS}/main

create:
	goose -dir ./migration create ${NAME} sql

up:
	goose -dir ./migration ${GOOSE_DRIVER} ${GOOSE_DBSTRING} up

down:
	goose -dir ./migration ${GOOSE_DRIVER} ${GOOSE_DBSTRING} down

status:
	goose -dir ./migration ${GOOSE_DRIVER} ${GOOSE_DBSTRING} status

test:
	go test -v ./api/user/service/action/__test__