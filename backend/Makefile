build:
	GOOS=linux go build -o app
	docker build -t gcr.io/neat-library-237522/say .
	rm -f app

push:
	gcloud docker -- push gcr.io/neat-library-237522/say
