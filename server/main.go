package main

import (
	pb "github.com/shivanik6z/greetservice/protos/greetpb"
)

type GreetingServer struct {
	pb.UnimplementedGreetingServiceServer
}

func main() {

}
