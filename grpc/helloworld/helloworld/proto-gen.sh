
SOURCE="./"
DEST="./"
FILE="*.proto"
gen(){
# docker run -dit --name protoc-go cbot918/protoc-go
  docker start protoc-go
  docker cp $SOURCE protoc-go:/app
  docker exec protoc-go protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $FILE
  docker cp protoc-go:/app/. $DEST
  docker exec protoc-go sh -c "rm ./*"
  ls $DEST
  # docker stop protoc-go && docker container rm protoc-go
}
gen