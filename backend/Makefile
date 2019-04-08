build:
	GOOS=linux go build -o app
	docker build -t ray/say .
	rm -f app