package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"oczko.com/game"

	"google.golang.org/grpc"

	pb "oczko.com/communication"
)

// Implements server

type SignalServer struct {
	game game.Game
}

func (s *SignalServer) MakeAction(ctx context.Context, request *pb.ActionRequest) (*pb.ActionReply, error) {
	fmt.Println(request.GetAction())
	return &pb.ActionReply{Winner: "Me"}, nil
}

func (s *SignalServer) CheckStatus(ctx context.Context, request *pb.StatusRequest) (*pb.StatusReply, error) {
	fmt.Println("Check status from ", request.PlayerID)
	var status string
	var statusType int
	status, statusType = s.game.GetStatus(int(request.PlayerID))
	return &pb.StatusReply{Status: status, Type: int32(statusType)}, nil
}

func (s *SignalServer) MakeMove(ctx context.Context, request *pb.MoveRequest) (*pb.MoveReply, error) {
	fmt.Println(request.PlayerID, " make move: ", request.Move)
	return &pb.MoveReply{Success: true, Message: s.game.MakeMove(int(request.PlayerID), request.Move)}, nil
}

func RunServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 33733))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterSignalServer(grpcServer, &SignalServer{game: game.NewGame()})
	grpcServer.Serve(lis)
}
