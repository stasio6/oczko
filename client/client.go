package client

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"google.golang.org/grpc"

	pb "oczko.com/communication"
)

// Implements client

func connectToServer() (*grpc.ClientConn, pb.SignalClient) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:33733", opts...)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	client := pb.NewSignalClient(conn)
	return conn, client
}

func callAction(client pb.SignalClient, msg string) {
	winner, err := client.MakeAction(context.Background(), &pb.ActionRequest{Action: "PLAY"})
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Response got: ", winner);
}

func checkStatus(client pb.SignalClient, id int) int {
	status, _ := client.CheckStatus(context.Background(), &pb.StatusRequest{PlayerID: int32(id)})
	fmt.Println("Status got: ", status.Status);
	return int(status.Type)
}

func makeMove(client pb.SignalClient, id int, move string) {
	message, _ := client.MakeMove(context.Background(), &pb.MoveRequest{PlayerID: int32(id), Move: move})
	fmt.Println("Message got: ", message.Message);
}

func MainClient() {
	conn, client := connectToServer()
	defer conn.Close()

	playerId := rand.Int()

	callAction(client, "Play")
	for {
		statusType := checkStatus(client, playerId)
		for statusType == 1 {
			time.Sleep(10 * time.Second)
			statusType = checkStatus(client, playerId)
		}
		if statusType == 0 {
			move := ""
			for move != "draw" && move != "pass"{
				fmt.Scanln(&move)
			}
			makeMove(client, playerId, move)
		} else if statusType == 2 {
			fmt.Println("Game ended!")
			break
		} else {
			panic("Error")
		}
	}
}