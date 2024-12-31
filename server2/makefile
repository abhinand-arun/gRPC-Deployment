proto:
	protoc --go_out=. --go-grpc_out=. ping.proto

run:
	go run main.go

deploy:
	kubectl apply -f grpc-server.yaml


docker-build:
	docker build -t abhinandarun/grpc-server:latest .

# Run the Docker container
docker-run:
	docker run -d -p 50051:50051 --name grpc-server abhinandarun/grpc-server:latest

# Stop the running Docker container
docker-stop:
	docker stop grpc-server

# Remove the stopped Docker container
docker-remove:
	docker rm grpc-server