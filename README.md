# Say-gRPC

Say-gRPC is project created with Go, gRPC and flite (Open source text to speech synthesis engine)

## Getting Started 

Install [GO](https://golang.org/) v1.12.3 on your machine, and udpate your `.bash_profile` with:

```bash 
export GOPATH="$HOME/go"
export GOBIN="$GOPATH/bin"
```

With `go` installed and `$GOPATH` set:
* `go get github.com/CheRayLiu/say-grpc`
* `cd $GOPATH/src/github.com/CheRayLiu/say-grpc`

### Running the application
* `cd say`
* `go run main.go -b 35.203.10.191:8080 "text"`, with "text" being the text desired to turn into speech
* ##### For macOS users:
    Run `afplay output.wav` to hear the speech

* ##### For users with a different OS :
    The sound file will generate in `$GOPATH/src/github.com/CheRayLiu/say-grpc/say` and can be played with a media player

## Running the application locally
* Download start [Docker](https://www.docker.com/)
* Run `make build` to build the project
* Run `make start`
* Open seperate terminal and `cd $GOPATH/src/github.com/CheRayLiu/say-grpc/say`
* Run `go run main.go text`, with "text" being the text desired to turn into speech

* ##### For macOS users:
    Run `afplay output.wav` to hear the speech

* ##### For users with a different OS :
    The sound file will be generated in `$GOPATH/src/github.com/CheRayLiu/say-grpc/say` and can be played with a media player

