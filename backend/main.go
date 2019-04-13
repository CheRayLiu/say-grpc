package main

import (
	"flag"
	"fmt"
	"net"
	"io/ioutil"
	"os/exec"

	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"github.com/Sirupsen/logrus"
	pb "github.com/CheRayLiu/say-grpc/api"
)

func main(){
	port := flag.Int("p", 8080, "port to listen to")
	flag.Parse()

	logrus.Infof("listening to port %d", port)
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
	f, err := ioutil.TempFile("","")
	if err != nil{
		return nil, fmt.Errorf("could not create temp file: %v", err)
	}

	if err := f.Close(); err != nil {
		return nil, fmt.Errorf("could not close temp file %s: %v", f.Name(), err)
	}

	cmd := exec.Command("flite", "-t", text.Text, "-o", f.Name())
	if data, err := cmd.CombinedOutput(); err != nil{
		return nil, fmt.Errorf("flite failed: %s", data) 
	}

	data, err := ioutil.ReadFile(f.Name())
	if err != nil{
		return nil, fmt.Errorf("could not read temp file %s: %v", f.Name(), err) 
	}
	return &pb.Speech{Audio: data}, nil
}
