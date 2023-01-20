DOCKFILE=./test/service/docker/docker-compose.yml

init-db:
	docker-compose up -d -f test/service/auth-docker-compose.yml

test-auth:
	go run test/cmd/testAuth/main.go