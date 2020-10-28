package main

import (
	"context"
	pb "course-info/grpc/proto"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	lis, err := net.Listen("tcp", ":8070")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterQuizServer(grpcServer, NewQuizServer())
	fmt.Printf("Server is running at :8070.\n")
	grpcServer.Serve(lis)
}

type QuizServer struct {
	signedUpUsers []*pb.User
	questionChan  chan *pb.Question
	pb.UnimplementedQuizServer
}

func NewQuizServer() *QuizServer {
	return &QuizServer{
		signedUpUsers: make([]*pb.User, 0),
		questionChan:  make(chan *pb.Question, 10),
	}
}

func (qs *QuizServer) Next(stream pb.Quiz_NextServer) error {
	for {
		question, err := stream.Recv()
		if err != nil {
			return status.Errorf(codes.NotFound, "Couldn't receive question from quiz master")
		}
		qs.questionChan <- question
	}
}

func (qs *QuizServer) Signup(user *pb.User, stream pb.Quiz_SignupServer) error {
	qs.signedUpUsers = append(qs.signedUpUsers, user)
	fmt.Println(qs.signedUpUsers)
	for {
		question := <-qs.questionChan
		err := stream.Send(question)
		if err != nil {
			return status.Errorf(codes.NotFound, "Couldn't send question to %s", user.GetUser())
		}
	}
}

func (qs *QuizServer) Vote(ctx context.Context, vote *pb.VoteRequest) (*pb.User, error) {
	fmt.Println(qs.signedUpUsers)
	for _, u := range qs.signedUpUsers {
		if vote.GetUser().User == u.User {
			return vote.GetUser(), nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "Couldn't find user %s", vote.GetUser())
}
