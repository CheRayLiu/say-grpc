package main

// import (
// 	"log"
// 	"os"
// 	"os/exec"
// )

import (
	"context"
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"github.com/Sirupsen/logrus"
	pb "github.com/CheRayLiu/say-grpc/api"
)

func main(){
	port := flag.Int("p", 8080, "port to listen to")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil{
		logrus.Fatalf("could not listent to port %d: %v", *port, err)
	}

	s := grpc.NewServer()
	pb.RegisterTextToSpeechServer(s, server{})
	err = s.Serve(lis)
	if err != nil{
		logrus.Fatalf("could not server: %v", err)
	}
}

type server struct{}

func (server) Say(ctx context.Context, text *pb.Text) (*pb.Speech, error){
	return nil, fmt.Errorf("not implemented")
}

	// cmd := exec.Command("flite", "-t", os.Args[1], "-o", "output.wav")
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// if err := cmd.Run(); err != nil{
	// 	log.Fatal(err) 
	// }