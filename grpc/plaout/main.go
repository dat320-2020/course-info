package main

import (
	"context"
	pb "course-info/grpc/proto"
	"flag"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	userName := flag.String("user", "", "the user name")
	flag.Parse()
	if *userName == "" {
		log.Fatal("user name is required")
	}

	conn, err := grpc.Dial(":8070", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewQuizClient(conn)
	user := &pb.User{User: *userName}
	// questionStream, err := client.Signup(context.Background(), user)
	_, err = client.Signup(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	// questionStream.

	vote := &pb.VoteRequest{
		QuestionId: 1,
		Vote:       1,
		User:       user,
	}
	winner, err := client.Vote(context.Background(), vote)
	if err != nil {
		fmt.Println("Got an error:", err)
	}
	fmt.Printf("The winner of round %d was %s\n", vote.QuestionId, winner.User)
}
