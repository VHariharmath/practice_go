package main

import (
	"log"
	"practice_go/grpc/demo-grpc/api"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect %s ", err)
	}

	defer conn.Close()

	c := api.NewPingClient(conn)

	response, err := c.SayHello(context.Background(), &api.PingMessage{Greeting: "foo"})

	if err != nil {
		log.Fatalf("Error when calling server %s ", err)
	}
	
	log.Printf("Response from server : %s ", response.Greeting);

}
