package main

import (
	"client/interconnect"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
)

const (
	defaultPort     = "5000"
	defaultHostname = "localhost"
)

var (
	port string
	host string
)

func init() {
	port = os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	host = os.Getenv("HOST")
	if host == "" {
		host = defaultHostname
	}
}

func readInputCharacter() string {
	fmt.Println("your character:")
	var s string
	fmt.Scan(&s)
	return string([]rune(s)[0])
}

func exitDefeat() {
	fmt.Println("DEFEAT")
	os.Exit(0)
}

func exitWin() {
	fmt.Println("WIN")
	os.Exit(0)
}

func printStatus(status *interconnect.WordStatus) {
	fmt.Println("----------")
	fmt.Printf("Your progress:\n%v\nfuel:%v\n", status.GameState, status.Fuel)
}

func GameStep(stream interconnect.Observer_ProcessGameClient) {
	status, err := stream.Recv()
	if err != nil {
		log.Fatalln(err)
	}

	printStatus(status)

	if status.GameStatus == interconnect.WordStatus_WIN {
		exitWin()
	}
	if status.GameStatus == interconnect.WordStatus_DEFEAT {
		exitDefeat()
	}

	stream.Send(&interconnect.WordGuess{
		Character: readInputCharacter(),
	})
}

func main() {
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := interconnect.NewObserverClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := client.ProcessGame(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		GameStep(stream)
	}
}
