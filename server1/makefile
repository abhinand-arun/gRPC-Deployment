proto:
	protoc --go_out=. --go-grpc_out=. ping.proto


deploy:
	kubectl apply -f k8s-deployment.yaml


docker-run:
	docker run -p 8080:8080 my-go-app

docker-build:
	docker build -t abhinandarun/my-go-app:latest .

docker-remove:
	docker run -p 8080:8080 my-go-app

