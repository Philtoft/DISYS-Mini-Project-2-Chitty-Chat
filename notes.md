[Inspiration](https://github.com/NaddiNadja/grpc101)

# Protoc command
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative time/time.proto

# Docker commands
- Build docker container: ``` docker build -t test . ```
- Run docker container ``` docker run -p 9080:9080 -tid test ```