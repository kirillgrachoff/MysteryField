package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"

	"github.com/kirillgrachoff/MysteryField/pkg/interconnect"
	"github.com/kirillgrachoff/MysteryField/pkg/util"
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
	port = util.GetOrDefault("PORT", defaultPort)

	host = util.GetOrDefault("HOST", defaultHostname)
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
	fmt.Printf("Your progress:\n%v\nfuel:%v\n\n", status.GameState, status.Fuel)
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

	err = stream.Send(&interconnect.WordGuess{
		Character: readInputCharacter(),
	})
	if err != nil {
		log.Fatalln(err)
	}
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
