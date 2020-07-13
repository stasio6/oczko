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

func joinGame(client pb.SignalClient, name string, id int) bool {
	reply, _ := client.JoinGame(context.Background(), &pb.JoinRequest{PlayerName: name, PlayerID: int32(id)})
	if reply.Success {
		fmt.Println("You successfully joined a new game!")
	} else {
		fmt.Println("An error occurred while joining the game")
	}
	return reply.Success
}

func checkStatus(client pb.SignalClient, id int) int {
	status, _ := client.CheckStatus(context.Background(), &pb.StatusRequest{PlayerID: int32(id)})
	fmt.Println("Status got: ", status.Status)
	return int(status.Type)
}

func makeMove(client pb.SignalClient, id int, move string) {
	message, _ := client.MakeMove(context.Background(), &pb.MoveRequest{PlayerID: int32(id), Move: move})
	fmt.Println("Message got: ", message.Message)
}

func playGame(client pb.SignalClient, name string) {
	playerId := rand.Int()
	if !joinGame(client, name, playerId) {
		fmt.Println("Couldn't join any game")
		return
	}
	for {
		statusType := 1
		for statusType == 1 {
			time.Sleep(time.Second)
			clearScreen()
			statusType = checkStatus(client, playerId)
		}
		if statusType == 0 {
			moves := []string{"draw", "pass"}
			move := openMenu(moves)
			makeMove(client, playerId, moves[move - 1])
		} else if statusType == 2 {
			fmt.Println("Game ended!")
			break
		} else {
			panic("Error")
		}
	}
}

func MainClient() {
	rand.Seed(int64(time.Now().Nanosecond()))

	conn, client := connectToServer()
	defer conn.Close()

	name := askForName(true)
	for {
		action := openMenu([]string{"Play", "Change name", "Exit"})

		switch action {
		case 1:
			playGame(client, name)
			break
		case 2:
			name = askForName(false)
			break
		case 3:
			return
		default:
			panic("Invalid action")
		}
	}
}