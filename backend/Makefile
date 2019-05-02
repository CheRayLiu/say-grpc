build:
	GOOS=linux go build -o app
	docker build -t ray/say-grpc .
	rm -f app

start: 
	docker run -p 8080:8080 ray/say-grpc

deploy: 
	GOOS=linux go build -o app
	docker build -t gcr.io/ray-say-grpc/say .
	rm -f app

push:
	gcloud docker -- push gcr.io/ray-say-grpc/say
