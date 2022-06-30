package main

import (
	pb "github.com/shivanik6z/greetservice/greet"
)

//pb "github.com/shivanik6z/greetservice/protos/greet"

type GreetingServer struct {
	pb.UnimplementedGreetingServiceServer
}

func main() {

}
