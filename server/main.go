package main

import (
	pb "github.com/shivanik6z/greetservice/protos"
)

type GreetingServer struct {
	pb.UnimplementedGreetingServiceServer
}

func main() {

}
