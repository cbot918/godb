DOCKFILE=./test/service/docker/auth-docker-compose.yml

init-db:
	docker-compose up -d -f $(DOCKFILE)

test-http:
	go run test/cmd/main.go