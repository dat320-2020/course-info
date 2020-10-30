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
	var (
		userName   = flag.String("user", "", "the user name")
		quizMaster = flag.Bool("master", false, "set this to run as quiz master")
	)
	flag.Parse()

	if !*quizMaster && *userName == "" {
		flag.Usage()
		return
	}

	conn, err := grpc.Dial(":8070", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewQuizClient(conn)
	if *quizMaster {
		quizMaxter(client)
	}

	user := &pb.User{User: *userName}
	questionStream, err := client.Signup(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}

	for {
		question, err := questionStream.Recv()
		if err != nil {
			fmt.Println("Got an error:", err)
		}
		fmt.Printf("New Question: %d\n-- %s", question.GetId(), question.GetQuestionText())
		for i, q := range question.GetAnswerText() {
			fmt.Printf("---- A%d: %s\n", i, q)
		}
		fmt.Print("What's your answer: ")
		var ansNum int32
		fmt.Scanf("%d", ansNum)
		vote := &pb.VoteRequest{
			QuestionId: 1,
			Vote:       ansNum,
			User:       user,
		}
		winner, err := client.Vote(context.Background(), vote)
		if err != nil {
			fmt.Println("Got an error:", err)
		}
		fmt.Printf("The winner of round %d was %s\n", vote.QuestionId, winner.User)
	}
}

func quizMaxter(client pb.QuizClient) {
	stream, err := client.Next(context.Background())
	if err != nil {
		fmt.Println("Got an error:", err)
	}
	x := &pb.Question{
		Id:           1,
		QuestionText: "Can we go home now?",
		AnswerText:   []string{"Not yet", "Soon", "Never", "Tomorrow", "Yes"},
	}
	err = stream.Send(x)
	if err != nil {
		fmt.Println("Got an error:", err)
	}

	questionTable := []*pb.Question{
		{
			Id:           1,
			QuestionText: "",
			AnswerText:   []string{"1a", "2a"},
		},
		{
			Id:           2,
			QuestionText: "Can we go home now?",
			AnswerText:   []string{"Not yet", "Soon", "Never", "Tomorrow", "Yes"},
		},
	}

	for _, q := range questionTable {
		fmt.Printf("Starting question round %d\n", q.GetId())
		err = stream.Send(x)
		if err != nil {
			fmt.Println("Got an error:", err)
		}
	}
}
