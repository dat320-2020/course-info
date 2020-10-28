package main

import (
	"bufio"
	"context"
	pb "course-info/grpc/proto"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"google.golang.org/grpc"
)

func main() {
	var (
		userName   = flag.String("user", "", "the user name")
		quizMaster = flag.Bool("master", false, "set if we want to be quiz master")
	)
	flag.Parse()

	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

	// if *userName == "" {
	// 	log.Fatal("user name is required")
	// }

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
		scanner.Scan()
		txt := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading text:", err)
			continue
		}
		ansNum, _ := strconv.Atoi(txt)
		vote := &pb.VoteRequest{
			QuestionId: 1,
			Vote:       int32(ansNum),
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

	// for i, q := range questionTable {
	// }
}

// var questionTable = []pb.Question{
// }
