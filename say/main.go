package mian

import (
	"flag"
	"log"
	"net"
	pb "github.com/CheRayLiu/say-grpc/api"
)
func main(){
	backend := flag.String("b", "localhost:8080", "address of the say backend")
	flag.Parse()
	output := flag.String("o", "output.wav", "wav file where the output will be written")
	conn, err := grpc.Dial(*backend)
	if err != nil{
		log.Fatalf("could not connection to %s: %v", *backend, err )
	}
	defer conn.close()

	client := pb.NewTextSpeechClient(conn)
	text : = &pb.Text{Text: "hello, there"}
	res, err := client.Say(context.Background(), text)
	if err != nil {
		log.Fatalf("could not say %s: %v", text.Text, err)
	}
	if err := ioutil.WriteFile(*output, res.Audio, 0666); err != nil{
		log.Fatalf("could not write to  %s: %v", *output, err)
	}

}