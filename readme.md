[Inspiration](https://github.com/NaddiNadja/grpc101)

# Protoc command
protoc 
    --go_out=. 
    --go_opt=paths=source_relative 
    --go-grpc_out=. 
    --go-grpc_opt=paths=source_relative time/time.proto

# Docker commands
- Build docker container: ``` docker build -t test . ```
- Run docker container ``` docker run -p 9080:9080 -tid test ```

[Tutorial to follow](https://www.youtube.com/watch?v=mML6GiOAM1w&ab_channel=TensorProgramming)

# Protoc command 2
protoc 
    --proto_path=service 
    --proto_path=. 
    --go_out=plugins=grpc:proto service.proto