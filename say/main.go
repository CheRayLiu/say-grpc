package main

import (
	"flag"
	"log"
	"google.golang.org/grpc"
	"io/ioutil"
	"context"
	"os"
	"fmt"
	
	pb "github.com/CheRayLiu/say-grpc/api"
)
func main(){
	backend := flag.String("b", "localhost:8080", "address of the say backend")
	flag.Parse()

	if flag.NArg() < 1{
		fmt.Printf("Usage:\n\t%s \"text to speech\"\n", os.Args[0])
		os.Exit(1)
	}
	output := flag.String("o", "output.wav", "wav file where the output will be written")
	conn, err := grpc.Dial(*backend, grpc.WithInsecure())
	if err != nil{
		log.Fatalf("could not connection to %s: %v", *backend, err )
	}
	defer conn.Close()

	client := pb.NewTextToSpeechClient(conn)
	text := &pb.Text{Text: flag.Arg(0)}
	res, err := client.Say(context.Background(), text)
	if err != nil {
		log.Fatalf("could not say %s: %v", text.Text, err)
	}
	if err := ioutil.WriteFile(*output, res.Audio, 0666); err != nil{
		log.Fatalf("could not write to  %s: %v", *output, err)
	}

}