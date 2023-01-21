#  protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld.proto


server:
	go run helloworld/greeter_server/main.go

client:
	go run helloworld/greeter_client/main.go


conn:
	/c/yaleProgram/ConEmuPack.221218/ConEmu64.exe -run bash ./test/conemu_auto.sh &


SOURCE = "pb/."
DEST = "pb/."
TEMP = helloworld.proto
proto:
# docker run -dit --name protoc-go cbot918/protoc-go
	docker exec protoc-go sh -c "rm /app/*"
	docker start protoc-go
	docker cp $(SOURCE) protoc-go:/app
	docker exec protoc-go protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $(TEMP)
	docker cp protoc-go:/app/. $(DEST)
	ls $(DEST)
# docker stop protoc-go && docker container rm protoc-go