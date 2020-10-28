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
	lis, err := net.Listen("tcp", "10.192.65.85:8070")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterQuizServer(grpcServer, NewQuizServer())
	fmt.Printf("Server is running at 10.192.65.85:8070.\n")
	grpcServer.Serve(lis)
}

type QuizServer struct {
	signedUpUsers []*pb.User
	pb.UnimplementedQuizServer
}

func NewQuizServer() *QuizServer {
	return &QuizServer{
		signedUpUsers: make([]*pb.User, 0),
	}
}

func (qs *QuizServer) Next(n pb.Quiz_NextServer) error {
	return nil
}

func (qs *QuizServer) Signup(user *pb.User, s pb.Quiz_SignupServer) error {
	qs.signedUpUsers = append(qs.signedUpUsers, user)
	fmt.Println(qs.signedUpUsers)
	return nil
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
